// --- Do not remove this imports!
import 'primeflex/primeflex.css';
import "primeicons/primeicons.css";
import "primereact/resources/themes/lara-light-indigo/theme.css";
// ---

import { Dialog } from 'primereact/dialog';
import { DataTable } from 'primereact/datatable';
import { Column } from 'primereact/column';
import { SchemasLogResponse } from '../../ApiClient';
import { Tag } from 'primereact/tag';
import { Knob } from 'primereact/knob';

type LogDetailDialogProps = {
    title: string;
    logData: SchemasLogResponse;
    visible: boolean;
    onHide: () => void;
};

export default function LogDetailDialog({ title, logData, visible, onHide }: LogDetailDialogProps) {
    const latestLog = logData;
    // const askToAi = () => {
    //     console.log(latestLog.exc_type);
    // }
    const getSeverity = (levelName: string) => {
        switch (levelName) {
            case 'INFO':
                return 'info';
            case 'WARNING':
                return 'warning';
            case 'ERROR':
                return 'danger';
            default:
                return null;
        }
    };

    function roundToFirstDecimal(num?: number): number {
        if (num === undefined) {
            return 0;
        }
        return Math.round(num * 10) / 10;
    }

    const levelNameTemplate = (log: SchemasLogResponse) => {
        console.log(log)
        return <Tag value={log.level_name} severity={getSeverity(log.level_name)}></Tag>;
    };

    const timestampTemplate = (log: SchemasLogResponse) => {
        // log.timestamp を Date オブジェクトに変換
        const date = new Date(log.timestamp);
        // 日時フォーマットのオプションを設定
        const options: Intl.DateTimeFormatOptions = {
            year: 'numeric',
            month: '2-digit',
            day: '2-digit',
            hour: '2-digit',
            minute: '2-digit',
            second: '2-digit',
            hour12: false // 24時間表記
        };
        // Intl.DateTimeFormat を使って日時をフォーマット
        const formattedDate = new Intl.DateTimeFormat('ja-JP', options).format(date);
        // フォーマットされた日時を返す
        return formattedDate;
    }

    const excDetailTemplate = (log: SchemasLogResponse) => {
        const excDetailFormatted = (log.exc_detail ?? "").split("\\n").join("\n");
        return (
            <pre className={`m-0 p-0 ${excDetailFormatted ? "p-3 bg-gray-800 text-gray-50 border-round-sm" : ""}`}>
                {excDetailFormatted}
            </pre>
        );
    }

    const cpuPercentTemplate = (log: SchemasLogResponse) => {
        return <Knob value={roundToFirstDecimal(log.cpu_percent)} readOnly />
    }

    const memoryPercentTemplate = (log: SchemasLogResponse) => {
        return <Knob value={roundToFirstDecimal(log.memory_percent)} readOnly />
    }

    return (
        <Dialog header={title} visible={visible} className="w-11 sm:w-9" onHide={onHide} modal>
            <div className="pt-4 sm:p-0">
                <div className="inline-flex py-1 px-2 align-items-center gap-2 border-round-sm surface-100">
                    <i className="pi pi-tag"></i>
                    <span className="font-semibold">{latestLog.system.category}</span>
                </div>
                <h2 className="m-0 py-3 px-4">{latestLog.system.name}</h2>

                <DataTable className="mb-4" header="ログ基本データ" value={[latestLog]}>
                    <Column field="level_name" header="Level" body={levelNameTemplate}/>
                    <Column field="message" header="Message" />
                    <Column field="file_name" header="File" />
                    <Column field="func_name" header="Function" />
                    <Column field="lineno" header="LineNo" />
                    <Column field="module" header="Module" />
                    <Column field="name" header="Name" />
                    <Column field="timestamp" header="Timestamp" body={timestampTemplate}/>
                </DataTable>

                <DataTable
                    className="mb-4"
                    header={
                        <div className="flex align-items-center justify-content-start">
                            エラー概要データ
                            {/* <Button icon="pi pi-comment" className="ml-2 p-button-rounded p-button-outlined ask-to-ai" onClick={askToAi} rounded text severity="secondary"/>
                            <Tooltip target={`.ask-to-ai`} content="Ask to Bard-AI." position="right"/> */}
                        </div>
                    }
                    value={[latestLog]}>
                    <Column field="exc_type" header="ExceptionType" />
                    <Column field="exc_value" header="ExceptionValue" />
                    <Column field="exc_detail" header="ExceptionDetail" body={excDetailTemplate}/>
                </DataTable>

                <DataTable className="mb-4" header="エラー詳細データ" value={latestLog.exc_traceback}>
                    <Column field="tb_filename" header="Filename" />
                    <Column field="tb_lineno" header="LineNo" />
                    <Column field="tb_name" header="Name" />
                    <Column field="tb_line" header="Line" />
                </DataTable>

                <DataTable header="オプションデータ" value={[latestLog]}>
                    <Column field="cpu_percent" header="CPU(%)" body={cpuPercentTemplate}/>
                    <Column field="process" header="Process" />
                    <Column field="process_name" header="ProcessName" />
                    <Column field="thread" header="Thread" />
                    <Column field="thread_name" header="ThreadName" />
                    <Column field="total_memory" header="TotalMemory" />
                    <Column field="available_memory" header="AvailableMemory" />
                    <Column field="memory_percent" header="Memory(%)" body={memoryPercentTemplate}/>
                    <Column field="used_memory" header="UsedMemory" />
                    <Column field="free_memory" header="FreeMemory" />
                    <Column field="cpu_user_time" header="UserTime(CPU)" />
                    <Column field="cpu_system_time" header="SystemTime(CPU)" />
                    <Column field="cpu_idle_time" header="IdleTime(CPU)" />
                    <Column field="levelno" header="LevelNumber" />
                    <Column field="system.name" header="SystemName" />
                </DataTable>
            </div>
        </Dialog>
    );
}
