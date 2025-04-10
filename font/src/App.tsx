import './App.css'

import message from "antd/es/message"
import ZxqLayout from "./pages/layout"
import IndexPage from "./pages/index"
import ProjectPage from "./pages/project"
import TaskPage from "./pages/task"
import AdminPage from "./pages/admin"
import AdminLayout from "./pages/admin/layout"
import CreateProject from "./pages/admin/createProject"
import Register from "./pages/admin/register"
import ResetRateLimit from "./pages/admin/resetRateLimit"
import UpdatePassword from "./pages/admin/updatePassword"
import SettingPage from "./pages/setting"
import MessageContext from "./context/message"


import {Route, Routes} from "react-router-dom";


function App() {

    const [middleMessageApi, middleMessageHolder] = message.useMessage({
        top: "30%",
        duration: 3,
        // @ts-ignore
        key: 'middle-message',
    })
    // @ts-ignore
    const [bottomRightMessageApi, bottomRightMessageHolder] = message.useMessage({
        top: "90%",
        // @ts-ignore
        right: "5%",
        duration: 3,
        key: 'bottom-right-message',
    })


    const messageContextValue = {
        middleApi: middleMessageApi,
        bottomRightApi: bottomRightMessageApi,

    }

  return (
    <MessageContext.Provider value={messageContextValue}>
        {middleMessageHolder}
        {bottomRightMessageHolder}
      <Routes>

        <Route index element={<IndexPage />} />
        <Route path={"/"} element={<ZxqLayout />}>
          <Route path="/project" >
            <Route index element={<ProjectPage />} />
          </Route>
          <Route path="/task"  >
            <Route index element={<TaskPage />} />
          </Route>
          <Route path="/admin" element={<AdminLayout />}>
              <Route index element={<AdminPage />} />
              <Route path={"create_project"} element={<CreateProject />} />
              <Route path={"register"} element={<Register />} />
              <Route path={"reset_rate_limit"}  element={<ResetRateLimit />} />
              <Route path={"update_password"} element={<UpdatePassword />} />
          </Route>
          <Route path="/setting" >
              <Route index element={<SettingPage />} />
          </Route>
        </Route>
    </Routes>
    </MessageContext.Provider>
  );
}

export default App
