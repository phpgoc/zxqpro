import { Button, Layout, Menu } from "antd";
import {
  HomeOutlined,
  SettingOutlined,
  SettingFilled,
  UsergroupAddOutlined, MessageOutlined
} from "@ant-design/icons";
import { Outlet, useLocation, useNavigate } from "react-router-dom";
import { BaseResponse, BaseResponseWithoutData } from "../types/response.ts";
import getRequestAndSetNavigateLocation from "../services/axios.ts";
import {useUserContext} from "../context/userInfo.tsx";
import { avatarUrl, isAdmin, serverUrl } from "../services/utils.ts";
import UserListSelect from "../components/userList.tsx";
import { useContext, useEffect, useState } from "react";
import MessageContext, { type MessageContextValue, SetMessageNumberContext } from "../context/message.tsx";
import { SSEMessage } from "../types/message.ts";

const { Header, Content } = Layout;

export default function ZxqLayout() {

  const [isSelectVisible, setIsSelectVisible] = useState(false);
  const [messageNumber, setMessageNumber] = useState(1);
  const items = [
    {
      key: "project",
      label: "Project",
      icon: <HomeOutlined style={{ fontSize: "5vh", lineHeight: "6vh" }} />,
    },
    {
      key: "task",
      label: "Task",
      icon: (
        <UsergroupAddOutlined style={{ fontSize: "5vh", lineHeight: "6vh" }} />
      ),
    },
    {
      key: "admin",
      label: "Admin",
      icon: <SettingFilled style={{ fontSize: "5vh", lineHeight: "6vh" }} />,
    },
    {
      key: "message",
      label: (
        <span style={{  alignItems: 'center' }}>
          Message
          {messageNumber > 0 && (
            <span
              style={{
                color: messageNumber > 0 ? '#ff4d4f' : 'inherit', // 红色警示色
              }}
            >
              ({messageNumber})
            </span>
          )|| ("(0)")}
        </span>
      ),
      icon: <MessageOutlined style={{ fontSize: "5vh", lineHeight: "6vh" }} />,
    },
    {
      key: "setting",
      label: "Setting",
      icon: <SettingOutlined style={{ fontSize: "5vh", lineHeight: "6vh" }} />,
    },
  ];
  const [sharedUserId, setSharedUserId] = useState(0);

  const navigate = useNavigate();
  const lct = useLocation();
  let request = getRequestAndSetNavigateLocation(navigate, lct);
  const currentPath = useLocation().pathname;

  const messageContext = useContext(MessageContext);
  const { middleMessageApi ,bottomRightMessageApi} = messageContext as MessageContextValue;
  const {user} = useUserContext()
  const avatarSrc =  avatarUrl(user.avatar)

  function logout() {
    request.post<BaseResponse>("user/logout").then((res) => {
      if (res.data.code == 0) {
        navigate("/");
      }
    });
  }

  function handleButtonClick() {
    setIsSelectVisible(true)
    setTimeout(() => {
      setIsSelectVisible(false)
    }, 10000)
  }

  function handleSelectChange(newUserId : number) {
    setSharedUserId(newUserId)
    // link: window.location.href,

    request.post<BaseResponseWithoutData>("message/share_link", {
      to_user_id: newUserId,
      link: location.pathname,
    }).then((res) => {
      if (res.data.code == 0) {
        middleMessageApi.success("分享成功").then();
      } else {
        middleMessageApi.error("分享失败").then();
      }
    });
    setIsSelectVisible(false)
  }

  useEffect(() => {
    if (!user || Object.keys(user).length === 0) {
      navigate("/");
    }else{
      if (!isAdmin(user.id)) {
        delete items[2];
      }
    }

    request.get("message/receive_list",{
      params: {
        page: 1,
        page_size: 1,
      },
    }).then((res) => {
      if (res.data.code == 0) {
        setMessageNumber(res.data.data.total);
      }
    })
    const eventSource = new EventSource(serverUrl() + "api/sse",{withCredentials:true});

    eventSource.onmessage = (event) => {
      const sseMessage = JSON.parse(event.data) as SSEMessage
      if(sseMessage.code !== 0) {
        middleMessageApi.error(sseMessage.message).then();
        if (sseMessage.code === 401) {
          navigate("/");
        }
        return;
      }
      setMessageNumber((prev) => (prev + 1))
      bottomRightMessageApi.success({
        content:  <>
          {sseMessage.message}
          {sseMessage.link ? (
            <>
              <br />
              &nbsp;{/* 空格 */}
              <a href={sseMessage.link!} onClick={(e) => {
                e.preventDefault(); // 阻止默认跳转
                window.location.href = sseMessage.link!; // 直接跳转
              }}>
                点击前往
              </a>
            </>
          ) : null}
        </>,
        duration: 30
      }).then()
    };
    return () => {
      eventSource.close();
    };
  }, []);

  return (
    <Layout>
      <Header
        style={{
          // 导航栏背景：渐变蓝色（从 #1890FF 到 #40a9ff）
          background: "linear-gradient(90deg, #1890FF 0%, #40a9ff 100%)",
          // 增加阴影提升层次感
          boxShadow: "0 2px 4px rgba(0, 0, 0, 0.08)",
          // height: '56px', // 适当增加高度让导航更舒展
          height: "7vh",
          display: "flex",
          justifyContent: "space-between",
          alignItems: "center",
          padding: "0 1vw", // 左右内边距，让内容不贴边
        }}
      >
        <div
          style={{
            height: "100%",
            width: "80%", // 左边元素占据50%宽度
            background: "linear-gradient(90deg, #1890FF 0%, #40a9ff 100%)",
          }}
        >
          <Menu
            style={{ height: "100%" }}
            mode="horizontal"
            items={items}
            selectedKeys={
              currentPath.includes("/project")
                ? ["project"]
                : currentPath.includes("/task")
                  ? ["task"]
                  : currentPath.includes("/admin")
                    ? ["admin"]
                    : currentPath.includes("/message")
                      ? ["message"]
                    : currentPath.includes("/setting")
                      ? ["setting"]
                      : []
            }
            onClick={(e) => {
              if (e.key === "project") navigate("/project");
              else if (e.key === "task") navigate("/task");
              else if (e.key === "message") navigate("/message");
              else if (e.key === "admin") navigate("/admin");
              else if (e.key === "setting") navigate("/setting");
            }}
          />
        </div>
        <div
          style={{
            display: "flex",
            alignItems: "center",
            height: "100%",
          }}
        >
          <div style={{ position: 'relative' }}>
              <Button
                style={{
                  position: 'relative',
                  zIndex: 1,
                }}
                onClick={handleButtonClick}
              >
                Share Link
              </Button>

            {isSelectVisible && (
              <div
                style={{
                  position: 'absolute',
                  top: 0,
                  left: 0,
                  zIndex: 2
                }}
              >
                <UserListSelect
                  userId={sharedUserId}
                  onChange={handleSelectChange}
                  includeAdmin={true}
                  filterSelf={true}
                />
              </div>
            )}
          </div>

          <img
            src={avatarSrc}
            alt="User Avatar"
            width={40}
            height={40}
            style={{ borderRadius: 10 }}
          />
          <span style={{ marginLeft: 10 }}>{user.user_name}</span>

          <Button
            type="primary"
            size="middle"
            style={{
              backgroundColor: "#fff", // 按钮背景为白色
              color: "#1890FF", // 按钮文字为蓝色主色
              border: "none",
              fontWeight: 600,
              padding: "8px 16px",
              borderRadius: 24, // 圆角更柔和
            }}
            onClick={logout}
          >
            Logout
          </Button>

        </div>
      </Header>
      <Content style={{ padding: "20px 24px 0" }}>

        <SetMessageNumberContext value={setMessageNumber}>
          <Outlet />
        </SetMessageNumberContext>
      </Content>
    </Layout>
  );
}
