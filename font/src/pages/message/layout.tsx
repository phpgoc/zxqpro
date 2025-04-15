import { Layout, Menu } from "antd";
import { useNavigate, useLocation, Outlet } from "react-router-dom";
import {
  HomeFilled,
  JavaScriptOutlined,
  PythonOutlined
} from "@ant-design/icons";

const { Sider, Content } = Layout;

export default function  MessageLayout  () {
  const navigate = useNavigate();
  const location = useLocation();

  const menuItems = [
    {
      key: "/message",
      icon: <HomeFilled />,
      label: "Received",
      path: "/message",
    },
    {
      key: "/message/send_list",
      icon: <JavaScriptOutlined />,
      label: "Send",
      path: "/message/send_list",
    },
    {
      key: "/message/manual_send",
      icon: <PythonOutlined />,
      label: "Manual Send",
      path: "/message/manual_send",
    },
  ];

  return (
    <Layout style={{ minHeight: "80vh" }}>
      <Sider width={200} style={{ background: "#fff" }}>
        <Menu
          mode="inline"
          selectedKeys={[location.pathname]}
          items={menuItems.map((item) => ({
            ...item,
            onClick: () => navigate(item.path),
          }))}
        />
      </Sider>
      <Layout>
        <Content
          style={{
            padding: 24,
            margin: 0,
            minHeight: 280,
          }}
        >
          <Outlet />
        </Content>
      </Layout>
    </Layout>
  )
}
