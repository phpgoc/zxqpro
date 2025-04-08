import {
    type RouteConfig,
    route,
    index,
    layout,
    prefix,
} from "@react-router/dev/routes";

export default [
    index("./pages/index.tsx"),

    layout("./pages/layout.tsx", [
            ...prefix("project",[
                route("/", "./pages/project/index.tsx"),
            ]),
        ...prefix("task",[
            route("/", "./pages/task/index.tsx"),
        ]),
        ...prefix("admin",[
            route("/", "./pages/admin/index.tsx"),
        ]),
        ...prefix("setting",[
            route("/", "./pages/setting/index.tsx"),
        ]),

    ]),

] satisfies RouteConfig;
