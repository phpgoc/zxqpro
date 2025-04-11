import { Form, Input, Checkbox, Button, Space } from "antd";
import getRequestAndSetNavigate from "../services/axios";
import { type BaseResponseWithoutData } from "../types/response";
import MessageContext, {
  type MessageContextValue,
} from "../context/message.tsx";
import { useContext } from "react";
import { useLocation, useNavigate } from "react-router-dom";
import { useUserContext } from "../context/userInfo.tsx";

interface LoginForm {
  name: string;
  password: string;
  use_mobile: boolean;
}

const LoginPage = () => {
  const [form] = Form.useForm<LoginForm>();
  const navigate = useNavigate();

  let request = getRequestAndSetNavigate(navigate, useLocation());

  const messageContext = useContext(MessageContext);
  const { middleApi } = messageContext as MessageContextValue;
  const {updateUser} = useUserContext()

  const submit = async (values: LoginForm) => {
    console.log("Received values of form: ", values);
    request
      .post<BaseResponseWithoutData>("user/login", JSON.stringify(values))
      .then((res) => {
        if (res.data.code === 0) {
          request.get("/user/info").then((info) => {
            updateUser(info.data.data);
            middleApi
              .success({
                content: "登录成功",
                duration: 1,
              })
              .then(() => {
                let redirectUrl = sessionStorage.getItem("redirectUrl");
                if (!redirectUrl || redirectUrl == "/") {
                  redirectUrl = "/project";
                }
                navigate(redirectUrl);
              });
          });
        } else {
          middleApi.error({
            content: res.data.message,
          });
        }
        console.log("Response from server: ", res.data.code);
      });
  };

  const onReset = () => {
    form.resetFields();
  };

  return (
    <div style={{ maxWidth: 400, margin: "0 auto", padding: 24 }}>
      <h2 style={{ textAlign: "center", marginBottom: 32 }}>登录</h2>

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
          rules={[{ required: true, message: "请输入用户名" }]}
        >
          <Input placeholder="请输入用户名" />
        </Form.Item>

        <Form.Item
          name="password"
          label="密码"
          rules={[{ required: true, message: "请输入密码" }]}
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
