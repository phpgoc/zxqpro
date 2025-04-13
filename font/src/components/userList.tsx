import { Select } from "antd";
import { UserInfo } from "../types/response.ts";
import { useEffect, useState } from "react";
import { useLocation, useNavigate } from "react-router-dom";
import getRequestAndSetNavigate from "../services/axios.ts";
import { useUserContext } from "../context/userInfo.tsx";

export default function UserListSelect({
  userId,
  onChange,
  projectId = 0,
  includeAdmin = false,
  filterSelf = false,
}: {
  userId: number;
  onChange: (newUserId: number) => void;
  projectId?: number;
  includeAdmin?: boolean;
  filterSelf?: boolean;
}) {
  const navigate = useNavigate();
  let request = getRequestAndSetNavigate(navigate, useLocation());
  const [userList, setUserList] = useState<UserInfo[]>([]);
  const {user} = useUserContext()
  useEffect(() => {
    let url = `user/list?project_id=${projectId}`
    if (includeAdmin) {
      url += "&include_admin=1";
    }
    request
      .get(url)
      .then((res) => {
        if (res.data.code === 0) {
          // Filter out the current user if filterSelf is true
          let filteredList = res.data.data.list as UserInfo[]
          if(filterSelf){
             filteredList = filteredList.filter(
              (u) => u.id != user.id
            );
          }
          setUserList(filteredList);
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


