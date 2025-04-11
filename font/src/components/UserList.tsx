import { Select } from "antd";
import { UserInfo } from "../types/response.ts";
import { useEffect, useState } from "react";
import { useLocation, useNavigate } from "react-router-dom";
import getRequestAndSetNavigate from "../services/axios.ts";

export default function UserListSelect({
  userId,
  onChange,
  projectId = 0,
}: {
  userId: number;
  onChange: (newUserId: number) => void;
  projectId?: number;
}) {
  const navigate = useNavigate();
  const lct = useLocation();
  let request = getRequestAndSetNavigate(navigate, lct);
  const [userList, setUserList] = useState<UserInfo[]>([]);
  useEffect(() => {
    request
      .get(`user/list?project_id=${projectId}`)
      .then((res) => {
        if (res.data.code === 0) {
          setUserList(res.data.data.list as UserInfo[]);
        }
      })
      .catch((error) => {
        console.error("Error fetching user list:", error);
      });
  }, []);
  return (
    <Select
      placeholder="Select User"
      onChange={onChange}
      onClear={() => onChange(0)}
      style={{ width: 200 }}
    >
      {userList.map((user) => (
        <Select.Option
          key={user.id}
          value={user.id}
          selected={userId === user.id}
        >
          {user.user_name}
        </Select.Option>
      ))}
    </Select>
  );
}
