// --- Do not remove this imports!
import 'primeflex/primeflex.css';
import "primeicons/primeicons.css";
import "primereact/resources/themes/lara-light-indigo/theme.css";
// ---

import { DataTable, DataTableSelectEvent } from 'primereact/datatable';
import { Column } from 'primereact/column';
import { SchemasLogResponse, SchemasPaginatedLogResponse } from "../../ApiClient";
import { Button } from 'primereact/button';
import LogDetailDialog from '../dialogs/LogDetailDialog';
import { useState } from 'react';
import { Card } from 'primereact/card';
import { Tag } from 'primereact/tag';

export default function LogDataTable({logData}: { logData: SchemasPaginatedLogResponse }) {
    const [logDetailVisible, setLogDetailVisible] = useState(false);
    const [selectedLog, setSelectedLog] = useState<SchemasLogResponse | null>(null);
    const getSeverity = (levelName: string) => {
        switch (levelName) {
            case 'INFO':
                return 'info';
            case 'WARNING':
                return 'warning';
            case 'ERROR':
                return 'danger';
            default:
                return undefined;
        }
    };

    const levelNameTemplate = (log: SchemasLogResponse) => {
        // console.log(log)
        return (
            // <Badge severity={getSeverity(log.level_name)} className="p-0" />
            <Tag value={log.level_name} severity={getSeverity(log.level_name)} className="py-0" />
        );
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

    const paginatorLeft = <Button type="button" icon="pi pi-refresh" text rounded severity="secondary" />;
    const paginatorRight = <Button type="button" icon="pi pi-download" text rounded severity="secondary" />;

    const onRowSelect = (event: DataTableSelectEvent) => {
        // console.log(event.data);
        setSelectedLog(event.data);
        setLogDetailVisible(true);
    };

    return (
        <Card>
            <DataTable size="small" paginator rows={25} rowsPerPageOptions={[5, 10, 25, 50]}
                value={logData.items}
                paginatorTemplate="RowsPerPageDropdown FirstPageLink PrevPageLink CurrentPageReport NextPageLink LastPageLink"
                currentPageReportTemplate="{first}~{last} / {totalRecords}"
                paginatorLeft={paginatorLeft} paginatorRight={paginatorRight}
                dataKey="id" filterDisplay="menu"
                globalFilterFields={['name', 'country.name', 'representative.name', 'status']}
                emptyMessage="No customers found."
                style={{ width: '100%' }} selectionMode="single"
                selection={selectedLog!} onRowSelect={onRowSelect}
                sortField="timestamp" sortOrder={-1} // ここに追加
            >
                <Column sortable filter field="level_name" header="" body={levelNameTemplate} style={{ minWidth: '4rem' }}/>
                <Column sortable field="timestamp" header="Timestamp" body={timestampTemplate} style={{ minWidth: '12rem' }}/>
                <Column filter field="message" header="Message" style={{ minWidth: '12rem' }}/>
                <Column filter field="name" header="Name" style={{ minWidth: '12rem' }}/>
                <Column filter field="func_name" header="Function" style={{ minWidth: '8rem' }} />
                <Column sortable filter dataType="numeric" field="lineno" header="LineNo" style={{ minWidth: '4rem' }} />
                <Column filter field="module" header="Module" style={{ minWidth: '4rem' }} />
                <Column filter field="exc_type" header="ExceptionType" style={{ minWidth: '8rem' }} />
                <Column sortable filter dataType="numeric" field="cpu_percent" header="CPU(%)" style={{ minWidth: '8rem' }} />
                <Column sortable filter dataType="numeric" field="memory_percent" header="Memory(%)" style={{ minWidth: '8rem' }} />
                {/* <Column field="system.name" header="System" style={{ minWidth: '12rem' }}/> */}
            </DataTable>
            {selectedLog != null &&
                <LogDetailDialog
                    title="最新ログ詳細"
                    logData={selectedLog}
                    visible={logDetailVisible}
                    onHide={() => {setLogDetailVisible(false); setSelectedLog(null);}}
                />
            }
        </Card>
    );
}
