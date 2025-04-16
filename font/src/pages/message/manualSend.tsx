import { Button, Input, Space } from "antd";
import UserMultipleSelect from "../../components/userMultipleSelect.tsx";
import { useContext, useState } from "react";
import { useLocation, useNavigate } from "react-router-dom";
import getRequestAndSetNavigateLocation from "../../services/axios.ts";
import MessageContext, { type MessageContextValue, SetMessageNumberContext } from "../../context/message.tsx";

export default function MessageManualSend() {

  const divStyle = {
    display: "flex",
    alignItems: "flex-start",
    width: "50%",
    marginBottom: 20,
  };
  const labelStyle = {
    width: "30%",
    // textAlign: "right",
    align : "right",
    marginRight: 10,
    marginTop: 5,
  }

  const [userIds, setUserSetUserIds] = useState([] as number[]);
  const [message, setMessage] = useState("");
  const [link, setLink] = useState("");

  const navigate = useNavigate();
  let request = getRequestAndSetNavigateLocation(navigate, useLocation());

  const messageContext = useContext(MessageContext);
  const { middleMessageApi } = messageContext as MessageContextValue;
  const setMessageNumber = useContext(SetMessageNumberContext);

  function handleSend() {
    request.post("message/manual", {
      user_ids: userIds,
      content: message,
      link: link==""?null:link,
    }).then((response) => {
      if (response.data.code == 0) {
        // middleMessageApi.success(response.data.message).then();
        setMessageNumber(prev => prev + 1)
        middleMessageApi.success(response.data.message).then()
      } else {
        // middleMessageApi.warning(response.data.message).then();
        middleMessageApi.error(response.data.message).then()
      }
    }).catch((error) => {
      console.error("Error sending message:", error);
    })

  }

  return (
    <>
      <Space direction="horizontal" style={{ marginBottom: 16 }}>
        <label >Send to</label><UserMultipleSelect onChange={setUserSetUserIds} filterSelf={true} includeAdmin={true} />
      </Space>

      <div
        style={{
          textAlign: "center",
          marginTop: 100,
          display: "flex",
          flexDirection: "column",
          alignItems: "center",
        }}
      >
        <div style={divStyle}>
          <label style={labelStyle}>
            Message
          </label>
          <Input
            value={message}
            onChange={(e) => setMessage(e.target.value)}
            placeholder="message"
            style={{ width: "70%" }}
          />
        </div>

        <div style={divStyle}>
          <label style={labelStyle}>
            Link
          </label>
          <Input
            value={link}
            onChange={(e) => setLink(e.target.value)}
            placeholder="link"
            style={{ width: "70%" }}
          />
        </div>

        <Space direction="horizontal" style={{ marginBottom: 16 }}>

        <Button
          onClick={handleSend}
          style={{ width: "15%", margin: "20px 1%" }}
        >
          Send
        </Button>

        <Button
          onClick={() => {
            setMessage("");
            setLink("");
            setUserSetUserIds([]);
          }}
          style={{ width: "15%", margin: "20px 1%" }}
        >
          Reset
        </Button>
        </Space>
      </div>
    </>

  )
}