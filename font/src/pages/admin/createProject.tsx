import { useContext, useState } from "react";
import getRequestAndSetNavigateLocation from "../../services/axios.ts";
import { useLocation, useNavigate } from "react-router-dom";
import { Input, Button, Form } from "antd";
import MessageContext, {
  type MessageContextValue,
} from "../../context/message.tsx";
import { BaseResponse } from "../../types/response.ts";
import * as React from "react";
import UserSelect from "../../components/userSelect.tsx";

export default function CreateProject() {
  const [name, setName] = useState("");
  const [description, setDescription] = useState("");
  const [ownerId, setOwnerId] = useState(0);

  const navigate = useNavigate();
  let request = getRequestAndSetNavigateLocation(navigate, useLocation());

  const messageContext = useContext(MessageContext);
  const { middleMessageApi } = messageContext as MessageContextValue;

  const handleSubmit = (_: React.FormEvent<HTMLFormElement>) => {
    // e.preventDefault();
    request
      .post<BaseResponse>("/admin/create_project", {
        name,
        description,
        owner_id: ownerId,
      })
      .then((response) => {
        if (response.data.code == 0) {
          middleMessageApi.success(response.data.message).then((_: any) => {
            setName("");
            setDescription("");
          });
        } else {
          middleMessageApi.warning(response.data.message).then();
        }
      })
      .catch((_) => {
        middleMessageApi.error("Create Project failed. Please try again.").then();
      });
  };

  return (
    <Form onFinish={handleSubmit}>
      <div
        style={{
          textAlign: "center",
          marginTop: 100,
          display: "flex",
          flexDirection: "column",
          alignItems: "center",
        }}
      >
        <div
          style={{
            display: "flex",
            alignItems: "center",
            width: "50%",
            marginBottom: 20,
          }}
        >
          <label style={{ width: "30%", textAlign: "right", marginRight: 10 }}>
            project name
          </label>
          <Input
            value={name}
            onChange={(e) => setName(e.target.value)}
            placeholder="Name"
            style={{ width: "70%" }}
          />
        </div>
        <div
          style={{
            display: "flex",
            alignItems: "flex-start",
            width: "50%",
            marginBottom: 20,
          }}
        >
          <label
            style={{
              width: "30%",
              textAlign: "right",
              marginRight: 10,
              marginTop: 5,
            }}
          >
            Description
          </label>
          <Input.TextArea
            value={description}
            onChange={(e) => setDescription(e.target.value)}
            placeholder="Description"
            style={{ width: "70%" }}
          />
        </div>
        <div
          style={{
            display: "flex",
            alignItems: "center",
            width: "50%",
            marginBottom: 20,
          }}
        >
          <label style={{ width: "30%", textAlign: "right", marginRight: 10 }}>
            Owner
          </label>
          <UserSelect userId={ownerId} onChange={setOwnerId} />
        </div>
        <div
          style={{ display: "flex", justifyContent: "center", width: "50%" }}
        >
          <Button
            type="primary"
            htmlType="submit"
            style={{ width: "15%", marginRight: 20 }}
          >
            Create
          </Button>
          <Button
            onClick={() => {
              setName("");
              setDescription("");
              setOwnerId(0);
            }}
            style={{ width: "15%", marginLeft: 20 }}
          >
            Reset
          </Button>
        </div>
      </div>
    </Form>
  );
}
