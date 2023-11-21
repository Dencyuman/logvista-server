// --- Do not remove this imports!
import 'primeflex/primeflex.css';
import "primeicons/primeicons.css";
import "primereact/resources/themes/lara-light-indigo/theme.css";
// ---

import { BarChart, Bar, XAxis, YAxis, CartesianGrid, Tooltip, Legend, ResponsiveContainer, Brush } from 'recharts';

export interface OverviewData {
    name: string;
    INFO: number;
    WARNING: number;
    ERROR: number;
}

export interface Layout {
    dateKey: string;
    stackId: string;
    fill: string;
}

// プロパティの型を定義
export interface OverviewChartProps {
    data: OverviewData[];
    layout: Layout[];
    customLayoutProps?: CustomLayouyProps;
    showBrush?: boolean;
    brushOrigin?: 'start' | 'end';
    brushLength?: number;
}

export interface CustomLayouyProps {
    width?: string;
    height?: string;
    top?: number;
    right?: number;
    left?: number;
    bottom?: number;
}

// 関数コンポーネントの引数は、propsオブジェクトです。
export function OverviewChart({
    data,
    layout,
    customLayoutProps = {
        width: '100%',
        height: '100%',
        top: 0,
        right: 0,
        left: 0,
        bottom: 0
    },
    showBrush = false,
    brushOrigin = 'end',
    brushLength = 30
}: OverviewChartProps) {
    const reversedData = [...data].reverse();

    let brushStart, brushEnd;
    if (brushOrigin === 'end') {
        brushStart = Math.max(reversedData.length - brushLength, 0);
        brushEnd = reversedData.length - 1;
    } else {
        brushStart = 0;
        brushEnd = Math.min(brushLength - 1, reversedData.length - 1);
    }

    return (
        <div style={{ width: customLayoutProps.width, height: customLayoutProps.height }}>
            <ResponsiveContainer width="100%" height="100%">
                <BarChart
                    data={reversedData}
                    margin={{
                        top: customLayoutProps.top,
                        right: customLayoutProps.right,
                        left: customLayoutProps.left,
                        bottom: customLayoutProps.bottom,
                    }}
                >
                    <CartesianGrid strokeDasharray="3 3" />
                    <XAxis dataKey="name" />
                    <YAxis />
                    <Tooltip />
                    <Legend
                        layout="horizontal"
                        wrapperStyle={{
                            display: 'flex',
                            justifyContent: 'center',
                            alignItems: 'center',
                            transform: 'translateX(30px)'
                        }}
                    />
                    {layout.map((layoutItem, index) => (
                        <Bar
                            key={index}
                            dataKey={layoutItem.dateKey}
                            stackId={layoutItem.stackId}
                            fill={layoutItem.fill}
                        />
                    ))}
                    {showBrush && (
                        <Brush
                            dataKey="name"
                            height={30}
                            stroke="#8884d8"
                            startIndex={brushStart}
                            endIndex={brushEnd}
                        />
                    )}
                </BarChart>
            </ResponsiveContainer>
        </div>
    );
}
