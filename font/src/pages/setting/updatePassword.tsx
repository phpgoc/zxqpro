import { useContext, useState } from "react";
import getRequestAndSetNavigate from "../../services/axios.ts";
import { useLocation, useNavigate } from "react-router-dom";
import { Input, Button } from "antd";
import MessageContext, {
  type MessageContextValue,
} from "../../context/message.tsx";
import { BaseResponse } from "../../types/response.ts";
import * as React from "react";

export default function UpdatePassword() {
  const [newPassowrd, setNewPassword] = useState("");
  const [oldPassword, setOldPassword] = useState("");
  const [newPassword2, setNewPassword2] = useState("");

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
      .post<BaseResponse>("/user/update_password", {  new_password: newPassowrd, old_password: oldPassword , new_password2: newPassword2})
      .then((response) => {
        if (response.data.code == 0) {
          middleMessageApi.success(response.data.message).then(
            (_: any) => {
              setNewPassword("");
              setOldPassword("");
              setNewPassword2("");
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
            Old Password
          </label>
          <Input.Password
            value={oldPassword}
            onChange={(e) => setOldPassword(e.target.value)}
            placeholder="Old Password"
            style={{ width: "70%" }}
          />
        </div>

        <div style={divStyle}>
          <label style={labelStyle}>
            New Password
          </label>
          <Input.Password
            value={newPassowrd}
            onChange={(e) => setNewPassword(e.target.value)}
            placeholder="new Password"
            style={{ width: "70%" }}
          />
        </div>

        <div style={divStyle}>
          <label style={labelStyle}>
            Confirm Password
          </label>
          <Input.Password
            value={newPassword2}
            onChange={(e) => setNewPassword2(e.target.value)}
            placeholder="Confirm Password"
            style={{ width: "70%" }}
          />
        </div>

        <Button
          type="primary"
          htmlType="submit"
          style={{ width: "15%", margin: "20px 1%" }}
        >
          Update
        </Button>

        <Button
          onClick={() => {
            setNewPassword("");
            setNewPassword2("");
            setOldPassword("");
          }}
          style={{ width: "15%", margin: "20px 1%" }}
        >
          Reset
        </Button>
      </div>
    </form>
  );
}
