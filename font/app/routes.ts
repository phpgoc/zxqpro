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
        route("/project", "./pages/project/index.tsx"),
        route("/task", "./pages/task/index.tsx"),
    ]),

    // ...prefix("concerts", [
    //     index("./concerts/home.tsx"),
    //     route(":city", "./concerts/city.tsx"),
    //     route("trending", "./concerts/trending.tsx"),
    // ]),
] satisfies RouteConfig;
