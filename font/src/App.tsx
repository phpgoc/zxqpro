import './App.css'

import message from "antd/es/message"
import ZxqLayout from "./pages/layout"
import IndexPage from "./pages/index"
import ProjectPage from "./pages/project"
import TaskPage from "./pages/task"
import AdminPage from "./pages/admin"
import SettingPage from "./pages/setting"
import MessageContext from "./context/message"


import {Route, Routes} from "react-router-dom";


function App() {
    const [middleMessageApi, middleMessageHolder] = message.useMessage({
        top: "30%",
        duration: 3,
    })
    const [bottomRightMessageApi, bottomRightMessageHolder] = message.useMessage({
        top: "90%",
        duration: 3,
        transitionName : "fade",
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
          <Route path="/admin" >
              <Route index element={<AdminPage />} />
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
