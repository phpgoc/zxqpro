import React from'react';
import { Layout, Menu } from 'antd';
import { HomeOutlined, UserOutlined, SettingOutlined } from '@ant-design/icons';

const { Header } = Layout;

const items = [
    {
        key: 'home',
        label: '首页',
        icon: <HomeOutlined />,
    },
    {
        key: 'users',
        label: '用户管理',
        icon: <UserOutlined />,
    },
    {
        key: 'settings',
        label: '设置',
        icon: <SettingOutlined />,
    },
];

const Login = () => {
    return (
        <Layout>
            <Header style={{ background: '#fff', padding: 0 }}>
                <Menu
                    mode="horizontal"
                    items={items}
                />
            </Header>
            {/* 其他页面内容 */}
        </Layout>
    );
};

export default Login;