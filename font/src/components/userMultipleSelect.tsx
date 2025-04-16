import { Select } from "antd";
import { UserInfo } from "../types/response.ts";
import { useEffect, useState } from "react";
import { useLocation, useNavigate } from "react-router-dom";
import getRequestAndSetNavigateLocation from "../services/axios.ts";
import { useUserContext } from "../context/userInfo.tsx";

export default function UserMultipleSelect({
  onChange,
  projectId = 0,
  includeAdmin = false,
  filterSelf = false,
}: {
  onChange: (newUserIds: number[]) => void;
  projectId?: number;
  includeAdmin?: boolean;
  filterSelf?: boolean;
}) {
  const navigate = useNavigate();
  let request = getRequestAndSetNavigateLocation(navigate, useLocation());
  const [options, setOptions] = useState<{ label: string; value: number }[]>([]);
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
          const newOptions = filteredList.map((user) => ({
            label: user.user_name,
            value: user.id,
          }))
          setOptions(newOptions)
        }
      })
      .catch((error) => {
        console.error("Error fetching user list:", error);
      });
  }, []);
  return (
      <Select
        mode="multiple"
        allowClear
        style={{ width: '40vw' }}
        placeholder="Select User"
        onChange={onChange}
        options={options}
      />
  );
}


