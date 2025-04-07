import { Outlet } from "react-router";
export default function Layout({ children }: { children: React.ReactNode }) {
    return (
        <div className="flex flex-col h-screen">
        <header className="bg-gray-800 text-white p-4">
            <a href="/project" className="ml-4" >Project</a>
            <a href="/task" className="ml-4">Task</a>
        </header>
        <main className="flex-grow p-4"><Outlet /></main>
        <footer className="bg-gray-800 text-white p-4 text-center">
            &copy; 2023 My Application
        </footer>
        </div>
    );
}