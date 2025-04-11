import { Layout, Menu } from "antd";
import { useNavigate, useLocation, Outlet } from "react-router-dom";
import {
  JavaOutlined,
  JavaScriptOutlined,
  PythonOutlined
} from "@ant-design/icons";

const { Sider, Content } = Layout;

const SettingLayout = () => {
  const navigate = useNavigate();
  const location = useLocation();

  const menuItems = [
    {
      key: "/setting",
      icon: <JavaOutlined />,
      label: "Setting Page",
      path: "/setting",
    },
    {
      key: "/setting/update_user",
      icon: <JavaScriptOutlined />,
      label: "Update Info",
      path: "/setting/update_user",
    },
    {
      key: "/setting/update_password",
      icon: <PythonOutlined />,
      label: "Update Password",
      path: "/setting/update_password",
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
  );
};

export default SettingLayout;
