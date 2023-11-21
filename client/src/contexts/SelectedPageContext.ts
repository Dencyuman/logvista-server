import { Dispatch, SetStateAction, createContext } from 'react';

export type Page = {
    name: string;
    path: string;
    iconClassName: string;
}

export type SelectedPageContextType = {
    selectedPage: Page;
    setSelectedPage: Dispatch<SetStateAction<Page>>;
}

export const SelectedPageContext = createContext<SelectedPageContextType | null>(null);

export const SelectedPageProvider = SelectedPageContext.Provider;