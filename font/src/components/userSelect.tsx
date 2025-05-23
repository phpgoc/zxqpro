import { Select } from "antd";
import { UserInfo } from "../types/response.ts";
import { useEffect, useState } from "react";
import { useLocation, useNavigate } from "react-router-dom";
import getRequestAndSetNavigateLocation from "../services/axios.ts";
import { useUserContext } from "../context/userInfo.tsx";

export default function UserSelect({
  userId,
  onChange,
  projectId = 0,
  includeAdmin = false,
  filterSelf = false,
                                     placeholder = "Select User"
}: {
  userId: number|null;
  onChange: (newUserId: number) => void;
  projectId?: number;
  includeAdmin?: boolean;
  filterSelf?: boolean;
  placeholder? : string;
}) {
  const navigate = useNavigate();
  let request = getRequestAndSetNavigateLocation(navigate, useLocation());
  const [userList, setUserList] = useState<UserInfo[]>([]);
  const {user} = useUserContext()
  useEffect(() => {
    let url = `project/user_list?id=${projectId}`
    if (projectId == 0){
      url = "user/list"
      if (includeAdmin) {
        url += "?include_admin=1";
      }
    }else if (includeAdmin) {
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
      placeholder={placeholder}
      onChange={
        onChange
      }
      value = {userId}
      // onClear={() => onChange(-1)}
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


