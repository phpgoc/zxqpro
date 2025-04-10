
import { Layout, Menu } from 'antd';
import { useNavigate, useLocation, Outlet } from 'react-router-dom';
import {
    HomeOutlined,
    PlusOutlined,
    UserAddOutlined,
    SyncOutlined,
    KeyOutlined
} from '@ant-design/icons';

const { Sider, Content } = Layout;

const AdminLayout = () => {
    const navigate = useNavigate();
    const location = useLocation();

    const menuItems = [
        {
            key: '/admin',
            icon: <HomeOutlined />,
            label: 'Admin Page',
            path: '/admin'
        },
        {
            key: '/admin/create_project',
            icon: <PlusOutlined />,
            label: 'Create Project',
            path: '/admin/create_project'
        },
        {
            key: '/admin/register',
            icon: <UserAddOutlined />,
            label: 'Register',
            path: '/admin/register'
        },
        {
            key: '/admin/reset_rate_limit',
            icon: <SyncOutlined />,
            label: 'Reset Rate Limit',
            path: '/admin/reset_rate_limit'
        },
        {
            key: '/admin/update_password',
            icon: <KeyOutlined />,
            label: 'Update Password',
            path: '/admin/update_password'
        }
    ];

    return (
        <Layout style={{ minHeight: '80vh' }}>
            <Sider width={200} style={{ background: '#fff' }}>
                <Menu
                    mode="inline"
                    selectedKeys={[location.pathname]}
                    items={menuItems.map(item => ({
                        ...item,
                        onClick: () => navigate(item.path)
                    }))}
                />
            </Sider>
            <Layout>
                <Content
                    style={{
                        padding: 24,
                        margin: 0,
                        minHeight: 280
                    }}
                >
                    {/* 嵌套路由的出口，显示匹配的子路由组件 */}
                    <Outlet />
                </Content>
            </Layout>
        </Layout>
    );
};

export default AdminLayout;
