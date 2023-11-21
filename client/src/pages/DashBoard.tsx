// --- Do not remove this imports!
import 'primeflex/primeflex.css';
import "primeicons/primeicons.css";
import "primereact/resources/themes/lara-light-indigo/theme.css";
// ---

import { useEffect, useState } from 'react';
import { useNavigate, useParams } from 'react-router-dom';
import apiClient, { SchemasSummary, SchemasSummaryData, SchemasSystemResponse, SchemasPaginatedLogResponse } from '../ApiClient';
import { overviewLayout } from '../service/SystemService';
import { OverviewData, OverviewChart } from '../components/charts/OverviewChart';
import LogDataTable from '../components/dashboard/LogDataTable';
import { Dropdown } from 'primereact/dropdown';
import { Menubar } from 'primereact/menubar';
import { SkeletonLogData } from '../components/dashboard/SkeletonLogData';
import { SkeletonChart } from '../components/dashboard/SkeletonChart';

const groupSystemsByCategory = (systems: SchemasSystemResponse[]) => {
    const grouped = systems.reduce((acc, system) => {
        // カテゴリーと名前がundefinedでないことを確認
        if (system.category && system.name) {
            acc[system.category] = acc[system.category] || [];
            acc[system.category].push({ label: system.name, value: system.name });
        }
        return acc;
    }, {} as Record<string, Array<{ label: string; value: string }>>);

    return Object.entries(grouped).map(([category, items]) => ({
        label: category,
        items
    }));
};

export default function DashBoard() {
    const navigate = useNavigate();
    const params = useParams();
    console.log(params);
    const [timeSpan] = useState<number>(10);
    const [dataCount] = useState<number>(1000);
    const [pageSize] = useState<number>(1200);
    const [logData, setLogData] = useState<SchemasPaginatedLogResponse>({
        items: [],
        limit: 0,
        page: 0,
        total: 0,
        total_pages: 0
    });
    const [summary, setSummary] = useState<SchemasSummary[]>([]);
    const [systems, setSystems] = useState<SchemasSystemResponse[]>([]);
    const [selectedSystemName, setSelectedSystemName] = useState<string | undefined>(params.systemName);
    const systemNameExists = selectedSystemName != undefined;
    const [isSummaryLoading, setIsSummaryLoading] = useState(false);
    const [isLogDataLoading, setIsLogDataLoading] = useState(false);

    useEffect(() => {
        const fetchSystemsData = async () => {
            try {
                const res = await apiClient.systemsGet();
                setSystems(res.data);
            } catch (error) {
                console.error("Error fetching systems data: ", error);
            }
        };

        fetchSystemsData();
    }, []);

    useEffect(() => {
        if (selectedSystemName && systems.length > 0) {
            setIsSummaryLoading(true);
            const fetchSummaryData = async () => {
                try {
                    const res = await apiClient.systemsSummaryGet(selectedSystemName, timeSpan, dataCount);
                    setSummary(res.data);
                } catch (error) {
                    console.error("Error fetching summary data: ", error);
                } finally {
                    setIsSummaryLoading(false);
                }
            };
            fetchSummaryData();
        }
    }, [selectedSystemName, systems, timeSpan, dataCount]);

    useEffect(() => {
        if (selectedSystemName && systems.length > 0) {
            setIsLogDataLoading(true);
            const fetchLogData = async () => {
                if (!selectedSystemName) return;

                try {
                    const res = await apiClient.logsGet(1, pageSize, undefined, undefined, undefined, selectedSystemName);
                    setLogData(res.data);
                } catch (error) {
                    console.error("Error fetching log data: ", error);
                } finally {
                    setIsLogDataLoading(false);
                }
            };

            fetchLogData();
        }
    }, [selectedSystemName, systems, pageSize]);

    useEffect(() => {
        setSelectedSystemName(params.systemName);
    }, [params.systemName]);

    function formatBaseTime(date: Date): string {
        return new Intl.DateTimeFormat('ja-JP', {
            hour: '2-digit',
            minute: '2-digit',
            second: '2-digit'
        }).format(date);
    }

    function convertToOverviewData(schemasData: SchemasSummaryData[] | undefined): OverviewData[] {
        if (!schemasData) {
            return [];
        }

        return schemasData.map((data) => ({
            name: formatBaseTime(new Date(data.base_time)),
            INFO: data.infolog_count,
            WARNING: data.warninglog_count,
            ERROR: data.errorlog_count
        }));
    }

    const systemOptions = groupSystemsByCategory(systems);

    const onSystemChange = (e: { value: string | undefined }) => {
        setSelectedSystemName(e.value);
        navigate(`/dashboard/${e.value}`);
    };

    const start = (
        <div className="flex align-items-center sm:pr-2">
            <img alt="logo" src="https://raw.githubusercontent.com/Dencyuman/logvista-cloud/main/client/src/assets/logo.png" height="40" className="mx-2"></img>
            <h2 className="hidden sm:block my-0">DashBoard</h2>
        </div>
    );

    const end = <Dropdown
        value={selectedSystemName}
        options={systemOptions}
        onChange={onSystemChange}
        placeholder="Select a System"
        className={"font-bold" + (selectedSystemName ? "" : " fadein animation-duration-1000 animation-iteration-infinite")}
        filter={true}
        optionGroupLabel="label"
        optionGroupChildren="items"
        style={{ minWidth: '15rem' }}
    />


// 通常のDashBoardコンポーネントのレンダリング
return (
    <div>
        <div className="flex flex-column align-items-start gap-3 py-2">
            <div className="w-full">
                <Menubar start={start} end={end} />
            </div>
            {systemNameExists && (
                <>
                    {isSummaryLoading || summary.length === 0 ? (
                        <SkeletonChart />
                    ) : (
                            <OverviewChart
                                data={convertToOverviewData(summary[0].data)}
                                layout={overviewLayout}
                                customLayoutProps={{
                                    width: "100%",
                                    height: "200px",
                                    top: 20,
                                    right: 80,
                                    left: 40,
                                    bottom: 5
                                }}
                                showBrush={true}
                                brushOrigin={'end'}
                                brushLength={60}
                            />
                        )
                    }
                    <div className="w-full">
                        {isLogDataLoading ? (
                            <SkeletonLogData />
                        ) : (
                            <LogDataTable logData={logData} />
                        )}
                    </div>
                </>
            )}
        </div>
    </div>
);

}