import { Form, Input, Checkbox, Button, Space } from 'antd';
import request from '../services/axios';
import { type BaseResponseWithoutData} from "../types/response";

interface LoginForm {
    name: string;
    password: string;
    use_mobile: boolean;
}

const LoginPage = () => {
    const [form] = Form.useForm<LoginForm>();


    const submit = async (values: LoginForm) => {
        console.log('Received values of form: ', values);
         request.post<BaseResponseWithoutData>('user/login', JSON.stringify(values)).then((res) => {
                if (res.data.code === 0) {
                   request.get("/user/info").then((info) => {
                     localStorage.setItem('userInfo', JSON.stringify(info.data));
                     window.location.href = '/project';
                   })
                }
           console.log('Response from server: ', res.data.code);
         })
    };

    const onReset = () => {
        form.resetFields();
    };

    return (
        <div style={{ maxWidth: 400, margin: '0 auto', padding: 24 }}>
            <h2 style={{ textAlign: 'center', marginBottom: 32 }}>登录</h2>

            <Form
                form={form}
                name="login-form"
                labelCol={{ span: 6 }}
                wrapperCol={{ span: 18 }}
                onFinish={submit}
            >
                <Form.Item
                    name="name"
                    label="用户名"
                    rules={[{ required: true, message: '请输入用户名' }]}
                >
                    <Input placeholder="请输入用户名" />
                </Form.Item>

                <Form.Item
                    name="password"
                    label="密码"
                    rules={[{ required: true, message: '请输入密码' }]}
                >
                    <Input.Password placeholder="请输入密码" />
                </Form.Item>

                <Form.Item name="use_mobile" valuePropName="checked" label="长期登录">
                    <Checkbox />
                </Form.Item>

                <Form.Item wrapperCol={{ offset: 6, span: 18 }}>
                    <Space>
                        <Button type="primary" htmlType="submit">
                            登录
                        </Button>
                        <Button onClick={onReset}>重置</Button>
                    </Space>
                </Form.Item>
            </Form>
        </div>
    );
};

export default LoginPage;