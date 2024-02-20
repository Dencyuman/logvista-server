// --- Do not remove this imports!
import 'primeflex/primeflex.css';
import "primeicons/primeicons.css";
import "primereact/resources/themes/lara-light-indigo/theme.css";
// ---

import AppTemplate from './templates/AppTemplate';
import { BrowserRouter, Routes, Route } from 'react-router-dom';
import DashBoard from './pages/DashBoard';
// import Setup from './pages/Setup';
// import Settings from './pages/Settings';
import Overview from './pages/Overview';

export default function App() {
    return (
        <BrowserRouter>
            <Routes>
                <Route path="" element={<AppTemplate />}>
                    <Route index path="" element={<Overview />} />
                    <Route path="dashboard/:systemName?" element={<DashBoard />} />
                    {/*<Route path="setup-guide" element={<Setup />} />*/}
                    {/*<Route path="settings" element={<Settings />} />*/}
                </Route>
            </Routes>
        </BrowserRouter>
    )
}

