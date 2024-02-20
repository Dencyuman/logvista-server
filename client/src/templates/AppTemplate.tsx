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
        { name: 'Overview', path: '/', iconClassName: 'pi pi-home'},
        { name: 'DashBoard', path: '/dashboard', iconClassName: 'pi pi-chart-bar' },
        // { name: 'SetupGuide', path: '/setup-guide', iconClassName: 'pi pi-box' },
        // { name: 'Settings', path: '/settings', iconClassName: 'pi pi-cog' },
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
                    <div className="bg-red-400 text-white py-2 px-4 mb-2 text-left border-round">
                        <p className="m-0">
                            これはデモページです。
                            <a href="https://takotakocyumans-organization.gitbook.io/logvista-docs/" target="_blank" className="text-white font-bold">
                                Logvista公式ドキュメント
                            </a>
                            を参照し、実環境はオンプレミスで構築してください。
                        </p>
                    </div>
                    <Outlet context={{handlePageChange}}/>
                </div>
            </div>
        </div>
    )
}