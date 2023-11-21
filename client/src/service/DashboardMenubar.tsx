import { Checkbox } from "primereact/checkbox";

const renderCheckboxItem = (label: string) => {
    return (
        <div>
            <Checkbox
                checked={false}
                onChange={(e) => console.log(e.checked)}
            />
            <span className="ml-2">{label}</span>
        </div>
    );
};

export const items = [
        {
            label: 'File',
            icon: 'pi pi-fw pi-file',
            items: [
                {
                    label: renderCheckboxItem('New'),
                    items: [
                        {
                            label: 'Bookmark',
                            icon: 'pi pi-fw pi-bookmark'
                        },
                        {
                            label: 'Video',
                            icon: 'pi pi-fw pi-video'
                        },

                    ]
                },
                {
                    label: 'Delete',
                    icon: 'pi pi-fw pi-trash'
                },
                {
                    separator: true
                },
                {
                    label: 'Export',
                    icon: 'pi pi-fw pi-external-link'
                }
            ]
        },
    ];
