import "./App.css";

import message from "antd/es/message";
import ZxqLayout from "./pages/layout";
import IndexPage from "./pages/index";
import SetHost from "./pages/setHost";
import TaskPage from "./pages/task";
import AdminPage from "./pages/admin";
import AdminLayout from "./pages/admin/layout";
import CreateProject from "./pages/admin/createProject";
import Register from "./pages/admin/register";
import ResetRateLimit from "./pages/admin/resetRateLimit";
import AdminUpdatePassword from "./pages/admin/updatePassword";
import MessageLayout from "./pages/message/layout";
import MessageIndex from "./pages/message";
import MessageSendList from "./pages/message/sendList";
import MessageManualSend from "./pages/message/manualSend";
import SettingLayout from "./pages/setting/layout";
import SettingPage from "./pages/setting";
import UpdateUser from "./pages/setting/updateUser";
import UpdatePassword from "./pages/setting/updatePassword";
import MessageContext from "./context/message";
import "@ant-design/v5-patch-for-react-19";
import { Route, Routes } from "react-router-dom";
import { UserProvider } from "./context/userInfo";
import ProjectIndex from "./pages/project";
import ProjectTasks from "./pages/project/tasks";
import ProjectView from "./pages/project/view";
import ProjectEdit from "./pages/project/edit";
import Help from "./pages/setting/help.tsx";
import { NotFound } from "./NotFound.tsx";

function App() {
  const [middleMessageApi, middleMessageHolder] = message.useMessage({
    top: "30%",
    duration: 3,
    // @ts-ignore
    key: "middle-message",
  });
  // @ts-ignore
  const [bottomRightMessageApi, bottomRightMessageHolder] = message.useMessage({
    top: "90%",
    duration: 10,
    // @ts-ignore
    key: "bottom-right-message",
  });

  const messageContextValue = {
    middleMessageApi: middleMessageApi,
    bottomRightMessageApi: bottomRightMessageApi,
  };

  return (
    <UserProvider>
      <MessageContext.Provider value={messageContextValue}>
        {middleMessageHolder}
        {bottomRightMessageHolder}
        <Routes>
          <Route path="*" element={<NotFound />} />
          <Route index element={<IndexPage />} />
          <Route path={"/set_host"} element={<SetHost />} />
          <Route path={"/"} element={<ZxqLayout />}>
            <Route path="/project">
              <Route index element={<ProjectIndex />} />
              <Route path=":id" element={<ProjectView />} />
              <Route path="tasks/:id" element={<ProjectTasks />} />
              <Route path="edit/:id" element={<ProjectEdit />} />
            </Route>
            <Route path="/task">
              <Route index element={<TaskPage />} />
            </Route>
            <Route path="/message" element={<MessageLayout />}>
              <Route index element={<MessageIndex />} />
              <Route path="send_list" element={<MessageSendList />} />
              <Route path="manual_send" element={<MessageManualSend />} />
            </Route>

            <Route path="/admin" element={<AdminLayout />}>
              <Route index element={<AdminPage />} />
              <Route path={"create_project"} element={<CreateProject />} />
              <Route path={"register"} element={<Register />} />
              <Route path={"reset_rate_limit"} element={<ResetRateLimit />} />
              <Route
                path={"update_password"}
                element={<AdminUpdatePassword />}
              />
            </Route>
            <Route path="/setting" element={<SettingLayout />}>
              <Route index element={<SettingPage />} />
              <Route path={"update_user"} element={<UpdateUser />} />
              <Route path={"update_password"} element={<UpdatePassword />} />
              <Route path={"help"} element={<Help />} />
            </Route>
          </Route>
        </Routes>
      </MessageContext.Provider>
    </UserProvider>
  );
}

export default App;
