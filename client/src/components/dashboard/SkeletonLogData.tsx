// --- Do not remove this imports!
import 'primeflex/primeflex.css';
import "primeicons/primeicons.css";
import "primereact/resources/themes/lara-light-indigo/theme.css";
// ---

import { Skeleton } from 'primereact/skeleton';

export function SkeletonLogData() {
    return (
        <div>
            <Skeleton shape="rectangle" className="w-full h-6rem mb-1" />
            {Array.from({ length: 10 }).map((_, index) => (
                <Skeleton key={index} shape="rectangle" className="w-full h-2rem mb-1" />
            ))}
            <Skeleton shape="rectangle" className="w-full h-4rem mb-1" />
        </div>
    );
}
