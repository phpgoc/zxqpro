import { useLocation, useNavigate } from "react-router-dom";
import getRequestAndSetNavigateLocation from "../services/axios.ts";
import React, { forwardRef, useEffect, useState } from "react";
import { UserInfo } from "../types/response.ts";
import { roleTypesMap } from "../types/project.ts";
import { Table } from "antd";

export interface UserInProjectMethods{
  addUser: (newUser: UserInfo) => void;
}

const   UserInProjectSelect = forwardRef<UserInProjectMethods,{id:number}> (( {id }, ref) =>{
  const navigate = useNavigate();
  let request = getRequestAndSetNavigateLocation(navigate, useLocation());
  const [userList, setUserList] = useState<UserInfo[]>([]);
  useEffect(() => {
    request
      .get(`user/list?project_id=${id}`)
      .then((res) => {
        if (res.data.code === 0) {
          const unsortedUserList = res.data.data.list as UserInfo[];
          const sortedUserList = unsortedUserList.sort((a, b) => a.role_type - b.role_type);
          setUserList(sortedUserList);
        }
      })
      .catch((error) => {
        console.error("Error fetching user list:", error);
      });
  }, []);

  const columns = [
    {
      title: 'User_name',
      dataIndex: 'user_name',
      key: 'name',
    },
    {
      title: 'Role',
      dataIndex: 'Role',
      key: 'Role',
      render: (_text: any, record: UserInfo) => {
        return roleTypesMap[record.role_type];
      },
    },
  ];

  const insertIntoSortedList = (list: UserInfo[], newItem: UserInfo): UserInfo[] => {
    let left = 0;
    let right = list.length;
    while (left < right) {
      const mid = Math.floor((left + right) / 2);
      if (list[mid].role_type < newItem.role_type) {
        left = mid + 1;
      } else {
        right = mid;
      }
    }
    const newList = [...list];
    newList.splice(left, 0, newItem);
    return newList;
  };

  // 添加新元素的函数

   const addUser = (newUser: UserInfo) => {
    setUserList((prevList) => insertIntoSortedList(prevList, newUser));
    console.log(newUser)
  };
   if (ref) {
     (ref as React.RefObject<{ addUser: (newUser: UserInfo) => void }>).current = {
       addUser,
     }
   }else{
      // console.error("ref is not defined.未必是错误，不想调用ref方法的就可以不传递ref");
   }

  return (
    <>
    <div
      style={{
        width: '80%', // 增大整个容器宽度
        margin: '0 auto', // 使容器居中
        backgroundColor: '#f0f0f0',
        padding: '1rem',
      }}
    >
      <h1
        style={{
          fontSize: '24px', // 增大表头字体大小
          textAlign: 'left',
          marginBottom: '1rem', // 增加表头与表格的间距
        }}
      >
        user in project
      </h1>
      <div
        style={{
          display: 'flex',
          justifyContent: 'flex-start', // 让表格左对齐
        }}
      >
        <Table
          dataSource={userList}
          columns={columns}
          rowKey={(item) => item.id}
          pagination={false} // 去掉分页
          style={{ fontSize: '18px', width: '100%' }} // 增大表格字体大小并占满容器宽度
        />
      </div>
    </div>

    </>
  );
})
export default UserInProjectSelect;