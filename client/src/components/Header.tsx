// --- Do not remove this imports!
import 'primeflex/primeflex.css';
import "primeicons/primeicons.css";
import "primereact/resources/themes/lara-light-indigo/theme.css";
// ---

import { Button } from 'primereact/button';
import { BreadCrumb } from 'primereact/breadcrumb';
import { useLocation, useNavigate } from 'react-router-dom';
import { Avatar } from 'primereact/avatar';
import { useEffect, useState } from 'react';

type BreadcrumbItem = {
    label: string;
    command?: () => void;
};

type HeaderProps = {
    setVisible: (value: boolean) => void;
    imageLink?: string;
};

export default function Header({ setVisible, imageLink }: HeaderProps) {
    const navigate = useNavigate();
    const location = useLocation();
    const [breadcrumbItems, setBreadcrumbItems] = useState<BreadcrumbItem[]>([]);

    useEffect(() => {
        const pathnames = location.pathname.split('/').filter(x => x);
        const breadcrumbItems = pathnames.map((path, index) => {
            const url = `/${pathnames.slice(0, index + 1).join('/')}`;
            return {
                label: path.charAt(0).toUpperCase() + path.slice(1),
                command: () => navigate(url),
            };
        });
        setBreadcrumbItems(breadcrumbItems);
    }, [location, navigate]);

    const home = {
        icon: 'pi pi-home',
        command: () => {
            navigate('/');
        }
    };

    return (
        <div className="flex justify-content-between">
            <div className="flex align-items-center">
                <Button icon="pi pi-bars" text rounded severity="secondary" aria-label="ナビゲーションメニュー" onClick={() => setVisible(true)}/>
                <BreadCrumb model={breadcrumbItems} home={home} className="border-none hidden sm:block" />
            </div>
            <div className="flex align-items-center">
                <Button icon="pi pi-search" text rounded className="p-button-rounded p-button-secondary" aria-label="検索" />
                <Button icon="pi pi-bell" text rounded className="p-button-rounded p-button-secondary" aria-label="通知" />
                <div className='pl-2 pr-3 cursor-pointer'>
                    {imageLink ? (
                        <Avatar image={imageLink} shape="circle" />
                    ) : (
                        <Avatar icon="pi pi-user" shape="circle" />
                    )}
                </div>
            </div>
        </div>
    )
}
