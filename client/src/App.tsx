// --- Do not remove this imports!
import 'primeflex/primeflex.css';
import "primeicons/primeicons.css";
import "primereact/resources/themes/lara-light-indigo/theme.css";
// ---

import AppTemplate from './templates/AppTemplate';
import { BrowserRouter, Routes, Route } from 'react-router-dom';
import Home from './pages/Home';
import DashBoard from './pages/DashBoard';
import Report from './pages/Report';
import Settings from './pages/Settings';
import Overview from './pages/Overview';

export default function App() {
    return (
        <BrowserRouter>
            <Routes>
                <Route path="" element={<AppTemplate />}>
                    <Route index path="" element={<Home />} />
                    <Route path="overview" element={<Overview />} />
                    <Route path="dashboard/:systemName?" element={<DashBoard />} />
                    <Route path="report" element={<Report />} />
                    <Route path="settings" element={<Settings />} />
                </Route>
            </Routes>
        </BrowserRouter>
    )
}

