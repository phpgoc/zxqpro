
import { useContext, useState } from "react";
import getRequestAndSetNavigate from "../../services/axios.ts";
import { useLocation, useNavigate } from "react-router-dom";
import { Input, Button } from "antd";
import MessageContext, {
  type MessageContextValue,
} from "../../context/message.tsx";
import { BaseResponse } from "../../types/response.ts";
import * as React from "react";
import UserListSelect from "../../components/userList.tsx";

export default function UpdatePassword() {
  const [userId , setUserId] = useState(0);
  const [password, setPassword] = useState("");

  const navigate = useNavigate();
  let request = getRequestAndSetNavigate(navigate, useLocation());

  const messageContext = useContext(MessageContext);
  const { middleMessageApi } = messageContext as MessageContextValue;

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

  const handleSubmit = (e: React.FormEvent<HTMLFormElement>) => {
    e.preventDefault();
    request
      .post<BaseResponse>("/admin/update_password", {  user_id: userId, password })
      .then((response) => {
        if (response.data.code == 0) {
          middleMessageApi.success(response.data.message).then(
            (_: any) => {
              setUserId(0)
              setPassword("");
            }
          );
        } else {
          middleMessageApi.warning(response.data.message).then();
        }
      })
      .catch((_) => {
        middleMessageApi.error("Registration failed. Please try again.").then();
      });
  };

  return (
    <form onSubmit={handleSubmit}>
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
            New Password
          </label>
          <Input.Password
            value={password}
            onChange={(e) => setPassword(e.target.value)}
            placeholder="New Password"
            style={{ width: "70%" }}
          />
        </div>

        <UserListSelect userId={userId} onChange={setUserId} />

        <Button
          type="primary"
          htmlType="submit"
          style={{ width: "15%", margin: "20px 1%" }}
        >
          Update
        </Button>

        <Button
          onClick={() => {
            setUserId(0)
            setPassword("");

          }}
          style={{ width: "15%", margin: "20px 1%" }}
        >
          Reset
        </Button>
      </div>
    </form>
  );
}
