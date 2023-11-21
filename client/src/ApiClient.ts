import { Configuration } from './api/configuration';
import { GeneralApiFactory, LogsApiFactory, MastersApiFactory, SystemsApiFactory } from './api/api';
import {
    SchemasErrorResponse,
    SchemasLog,
    SchemasLogResponse,
    SchemasPaginatedLogResponse,
    SchemasResponseMessage,
    SchemasSummary,
    SchemasSummaryData,
    SchemasSystemResponse,
    SchemasTraceback,
    SchemasTracebackResponse
} from './api';

export const VITE_API_URL = import.meta.env.VITE_API_URL || "http://localhost:8080"
export const BASE_PATH = import.meta.env.VITE_API_BASE_PATH || "/api/v1"
export const BASE_URL = VITE_API_URL + BASE_PATH

const generalApiClient = GeneralApiFactory(new Configuration({ basePath: BASE_URL }))
const logsApiClient = LogsApiFactory(new Configuration({ basePath: BASE_URL }))
const mastersApiClient = MastersApiFactory(new Configuration({ basePath: BASE_URL }))
const systemsApiClient = SystemsApiFactory(new Configuration({ basePath: BASE_URL }))

const apiClient = {
    ...generalApiClient,
    ...logsApiClient,
    ...mastersApiClient,
    ...systemsApiClient,
}

export type {
    SchemasErrorResponse,
    SchemasLog,
    SchemasLogResponse,
    SchemasPaginatedLogResponse,
    SchemasResponseMessage,
    SchemasSummary,
    SchemasSummaryData,
    SchemasSystemResponse,
    SchemasTraceback,
    SchemasTracebackResponse
}
export default apiClient