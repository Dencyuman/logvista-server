export const overviewLayout = [
    {
    dateKey: 'INFO',
    stackId: 'logCategory',
    fill: '#3B82F6',
    },
    {
    dateKey: 'WARNING',
    stackId: 'logCategory',
    fill: '#F59E0A',
    },
    {
    dateKey: 'ERROR',
    stackId: 'logCategory',
    fill: '#EF4443',
    },
];

export const SystemService = {
    getSystemsData() {
        function randomPastTime() {
            const now = new Date();
            const randomOffset = Math.floor(Math.random() * 14); // 0から14までのランダムな数
            let randomDate;
            switch (true) {
                case randomOffset < 7: // 数分前
                    randomDate = new Date(now.getTime() - randomOffset * 1000 * 60); // 0から6分前
                    break;
                case randomOffset < 12: // 数時間前
                    randomDate = new Date(now.getTime() - (randomOffset - 5) * 1000 * 60 * 60); // 1から5時間前
                    break;
                case randomOffset <= 14: // 数日前
                    randomDate = new Date(now.getTime() - (randomOffset - 10) * 1000 * 60 * 60 * 24); // 1から4日前
                    break;
            }
            return randomDate;
        }

        function generateSampleData() {
            const data = [];
            const currentTime = new Date();

            for (let i = 0; i < 12; i++) {
                // 現在の時刻からi時間前の時刻を計算
                const time = new Date(currentTime.getTime() - i * 60 * 60 * 1000);

                // 時刻をフォーマットする (例: '10:00 AM')
                const formattedTime = time.toLocaleTimeString('en-US', { hour: 'numeric', minute: 'numeric', hour12: true });

                // データオブジェクトを生成
                const dataEntry = {
                    name: formattedTime, // 時刻をnameとして使用
                    INFO: Math.round(Math.random() * 350),
                    WARNING: Math.round(Math.random() * 50),
                    ERROR: Math.round(Math.random() * 100),
                };

                // 配列の先頭に追加して、時刻が降順になるようにする
                data.unshift(dataEntry);
            }

            return data;
        }

        function getRandomStatus() {
            const statuses = ['NORMAL', 'WARNING', 'ERRORED'];
            const randomIndex = Math.floor(Math.random() * statuses.length);
            return statuses[randomIndex];
        }

        return [
            {
                id: '1000',
                name: 'SnapCheck',
                latestTimestamp: randomPastTime(),
                status: getRandomStatus(),
                category: 'API Server',
                data: generateSampleData(),
            },
            {
                id: '1001',
                name: '小工程変更API',
                latestTimestamp: randomPastTime(),
                status: getRandomStatus(),
                category: 'API Server',
                data: generateSampleData()
            },
            {
                id: '1002',
                name: 'TGK Util API',
                latestTimestamp: randomPastTime(),
                status: getRandomStatus(),
                category: 'API Server',
                data: generateSampleData()
            },
            {
                id: '1003',
                name: 'PAS',
                latestTimestamp: randomPastTime(),
                status: getRandomStatus(),
                category: 'Automation',
                data: generateSampleData()
            },
            {
                id: '1004',
                name: 'PASK',
                latestTimestamp: randomPastTime(),
                status: getRandomStatus(),
                category: 'Automation',
                data: generateSampleData()
            },
            {
                id: '1005',
                name: '北九州メール転送自動化',
                latestTimestamp: randomPastTime(),
                status: getRandomStatus(),
                category: 'Automation',
                data: generateSampleData()
            },
            {
                id: '1006',
                name: 'ハンド検収入力自動化',
                latestTimestamp: randomPastTime(),
                status: getRandomStatus(),
                category: 'Automation',
                data: generateSampleData()
            },
            {
                id: '1007',
                name: 'ハンド売上実績入力自動化',
                latestTimestamp: randomPastTime(),
                status: getRandomStatus(),
                category: 'Automation',
                data: generateSampleData()
            },
            {
                id: '1008',
                name: 'Zaikan-福岡 IE',
                latestTimestamp: randomPastTime(),
                status: getRandomStatus(),
                category: 'Application',
                data: generateSampleData()
            },
            {
                id: '1009',
                name: 'Zaikan-北九州 IE',
                latestTimestamp: randomPastTime(),
                status: getRandomStatus(),
                category: 'Application',
                data: generateSampleData()
            },
            {
                id: '1010',
                name: '労働時間管理システム(45)',
                latestTimestamp: randomPastTime(),
                status: getRandomStatus(),
                category: 'Automation',
                data: generateSampleData()
            },
            {
                id: '1011',
                name: '労働時間管理システム(120)',
                latestTimestamp: randomPastTime(),
                status: getRandomStatus(),
                category: 'Automation',
                data: generateSampleData()
            },
        ];
    },

    getProductsMini() {
        return Promise.resolve(this.getSystemsData().slice(0, 5));
    },

    getProductsSmall() {
        return Promise.resolve(this.getSystemsData().slice(0, 10));
    },

    getSystems() {
        return Promise.resolve(this.getSystemsData());
    },
};

