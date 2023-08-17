import { createRouter, createWebHistory } from 'vue-router';
import AppBody from '../components/AppBody.vue';
import AppLogin from '../components/AppLogin.vue';
import AppSchemas from '../components/AppSchemas.vue';
import TableInfoComponent from '../components/TableInfoComponent.vue'; // Remember to create this component
import AppCustomQuery from '../components/AppCustomQuery.vue';
import AppViews from '../components/AppViews.vue';
import AppInfo from '../components/AppInfo.vue';
import AppStats from '../components/AppStats.vue';

const routes = [
    {
        path: "/",
        name: "Home",
        component: AppBody
    },
    {
        path: "/login",
        name: "Login",
        component: AppLogin
    },
    {
        path: "/schemas",
        name: "Schemas",
        component: AppSchemas
    },
    {
        path: '/schemas/:tableName',
        name: 'TableInfo',
        component: TableInfoComponent,
        props: true // this will set the tableName param as a prop on your component
    },
    {
        path: '/custom-query',
        name: 'CustomQuery',
        component: AppCustomQuery
    },
    {
        path: '/views',
        name: 'Views',
        component: AppViews
    },
    {
        path: '/info',
        name: 'Info',
        component: AppInfo
    },
    {
        path: '/db-stats',
        name: 'DBStats',
        component: AppStats
    }
];

const router = createRouter({history: createWebHistory(), routes});
export default router;
