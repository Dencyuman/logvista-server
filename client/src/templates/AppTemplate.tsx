// --- Do not remove this imports!
import 'primeflex/primeflex.css';
import "primeicons/primeicons.css";
import "primereact/resources/themes/lara-light-indigo/theme.css";
// ---

import { SetStateAction, useEffect, useMemo, useState } from 'react';

import Sidescreen from '../components/Sidescreen';
import Header from '../components/Header';
import { Outlet } from 'react-router-dom';

export type Page = {
    name: string;
    path: string;
    iconClassName: string;
}

export type AppContextType = {
    handlePageChange: (page: Page) => void;
};

export default function AppTemplate() {
    const [visible, setVisible] = useState(false);
    const handlePageChange = (page: SetStateAction<{ name: string; path: string; iconClassName: string; }>) => {
        setSelectedPage(page);
    };
    const pages = useMemo(() => ([
        { name: 'Home', path: '/', iconClassName: 'pi pi-home' },
        { name: 'Overview', path: '/overview', iconClassName: 'pi pi-th-large'},
        { name: 'DashBoard', path: '/dashboard', iconClassName: 'pi pi-chart-bar' },
        // { name: 'Report', path: '/report', iconClassName: 'pi pi-chart-pie' },
        { name: 'Settings', path: '/settings', iconClassName: 'pi pi-cog' },
    ]), []);

    const [selectedPage, setSelectedPage] = useState(pages[0]);
    const imageLink: undefined = undefined;

    useEffect(() => {
        const currentPath = window.location.pathname;
        const matchingPage = pages.find(page => page.path === currentPath);
        if (matchingPage) {
            setSelectedPage(matchingPage);
        }
    }, [pages]);

    return (
        <div className="card grid ">
            <div className="col-12">
                <Sidescreen visible={visible} setVisible={setVisible} selectedPage={selectedPage} setSelectedPage={setSelectedPage} pages={pages}/>
                <Header setVisible={setVisible} imageLink={imageLink ? imageLink : undefined}/>
                <div className="px-2 sm:px-6 m-auto container">
                    <Outlet context={{ handlePageChange }}/>
                </div>
            </div>
        </div>
    )
}