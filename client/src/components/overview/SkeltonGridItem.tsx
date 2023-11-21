// --- Do not remove this imports!
import 'primeflex/primeflex.css';
import "primeicons/primeicons.css";
import "primereact/resources/themes/lara-light-indigo/theme.css";
// ---

import { Skeleton } from "primereact/skeleton";

export function SkeletonGridItem() {
    return (
        <div className="col-12 py-3 sm:px-3 lg:col-6 xl:col-4">
            <Skeleton shape="rectangle" width="100%" height="400px" className="mb-2"></Skeleton>
            <Skeleton width="60%" className="mb-2"></Skeleton>
            <Skeleton width="80%" className="mb-2"></Skeleton>
            <Skeleton width="40%"></Skeleton>
        </div>
    )
}