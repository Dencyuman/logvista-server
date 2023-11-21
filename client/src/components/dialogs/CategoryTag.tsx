// --- Do not remove this imports!
import 'primeflex/primeflex.css';
import "primeicons/primeicons.css";
import "primereact/resources/themes/lara-light-indigo/theme.css";
// ---

import React, { useRef, useState } from 'react';
import apiClient, { SchemasSummary } from "../../ApiClient";
import { Dialog } from 'primereact/dialog';
import { AutoComplete } from 'primereact/autocomplete';
import { Button } from 'primereact/button';
import { Toast } from 'primereact/toast';

type CategoryTagProps = {
    summary: SchemasSummary;
    onCategoryChange: () => void;
    uniqueCategories: string[];
    fetchCategories: () => Promise<void>;
};

export default function CategoryTag({ summary, onCategoryChange, uniqueCategories, fetchCategories }: CategoryTagProps) {
    const [filteredCategories, setFilteredCategories] = useState<string[]>([]);
    const [selectedCategory, setSelectedCategory] = useState<string>(summary.category || "");
    const [displayDialog, setDisplayDialog] = useState(false);
    const toast = useRef<Toast>(null);

    const onConfirm = async () => {
        try {
            if (!summary.name) throw new Error("summary.name is undefined");
            await apiClient.systemsSystemNamePut(summary.name, { category: selectedCategory });
            setDisplayDialog(false);
            await fetchCategories();
            onCategoryChange();
        } catch (error) {
            if (toast.current) {
                toast.current.show({ severity: 'error', summary: 'エラー', detail: 'カテゴリの更新に失敗しました。', life: 3000});
            }
        }
    };

    const searchCategory = (event: { query: string }) => {
        const filtered = uniqueCategories.filter((category) =>
            category.toLowerCase().includes(event.query.toLowerCase())
        );
        setFilteredCategories(filtered);
    };

    const onClick = () => {
        setDisplayDialog(true);
    };

    const onHideDialog = () => {
        setDisplayDialog(false);
    };

    const handleKeyDown = (event: React.KeyboardEvent) => {
        // Ctrl + Enter が押されたときだけ onConfirm を実行する
        if (event.key === 'Enter' && event.ctrlKey) {
            onConfirm();
        }
    };

    return (
        <div>
            <Toast ref={toast} position="bottom-left" />
            <Button
                className="flex align-items-center gap-2 py-2 px-0 border-round-sm cursor-pointer hover:surface-100 p-button-text p-button-plain"
                onClick={onClick}
                severity="secondary"
            >
                <i className="pi pi-tag"></i>
                <span className="font-semibold">{summary.category}</span>
            </Button>

            <Dialog className='w-9 max-w-30rem py-0' visible={displayDialog} onHide={onHideDialog} header={`カテゴリ選択: ${summary.name}`}>
                <div className="w-full flex flex-column text-900">
                    <div className="grid">
                        <AutoComplete
                            value={selectedCategory}
                            suggestions={filteredCategories}
                            completeMethod={searchCategory}
                            dropdown
                            onChange={(e) => setSelectedCategory(e.value)}
                            className='my-2 col'
                            onKeyDown={handleKeyDown}
                        />
                    </div>
                    {selectedCategory !== "" && <p>カテゴリを <span className="font-semibold">{selectedCategory}</span> に設定しますか？</p>}
                    <div className="flex gap-2 justify-content-end">
                        <Button className="p-button" size='small' style={{ minWidth: '120px' }} label="Yes" onClick={onConfirm} />
                        <Button className="p-button-outlined" size='small' style={{ minWidth: '120px' }} label="Cancel" onClick={onHideDialog} />
                    </div>
                </div>
            </Dialog>
        </div>
    );
}
