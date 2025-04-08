import './App.css'
import ZxqLayout from "./pages/layout"
import IndexPage from "./pages/index"
import ProjectPage from "./pages/project"
import TaskPage from "./pages/task"
import AdminPage from "./pages/admin"
import SettingPage from "./pages/setting"


import {Route, Routes} from "react-router-dom";

function App() {

  return (
      <Routes>

        <Route index element={<IndexPage />} />
        <Route path={"/"} element={<ZxqLayout />}>
          <Route path="/project" >
            <Route index element={<ProjectPage />} />
          </Route>
          <Route path="/task"  >
            <Route index element={<TaskPage />} />
          </Route>
          <Route path="/admin" >
              <Route index element={<AdminPage />} />
          </Route>
          <Route path="/setting" >
              <Route index element={<SettingPage />} />
          </Route>
        </Route>
    </Routes>
  );
}

export default App
