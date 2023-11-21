// --- Do not remove this imports!
import 'primeflex/primeflex.css';
import "primeicons/primeicons.css";
import "primereact/resources/themes/lara-light-indigo/theme.css";
// ---

import { Skeleton } from 'primereact/skeleton';

export function SkeletonChart() {
    return (
        <div className="flex flex-column w-full align-items-center py-3 gap-1">
            <Skeleton className="w-8 sm:w-10" height="120px" />
            <Skeleton className="w-8 sm:w-10" height="30px" />
            <Skeleton className="w-7 sm:w-4 xl:w-2" height="15px" />
        </div>
    );
}
