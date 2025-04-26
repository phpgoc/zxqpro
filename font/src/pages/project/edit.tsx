import UserInProjectSelect, { UserInProjectMethods } from "../../components/userInProjectSelect.tsx";
import { useLocation, useNavigate, useParams } from "react-router-dom";
import { useContext, useEffect, useRef, useState } from "react";
import { BaseResponse, ProjectInfo, UserInfo } from "../../types/response.ts";
import getRequestAndSetNavigateLocation from "../../services/axios.ts";
import { Button, Checkbox, Input, Select } from "antd";
import { RoleTypeForAddSelect } from "../../components/roleType.tsx";
import MessageContext, { type MessageContextValue } from "../../context/message.tsx";
import { parseIdToNumber } from "../../services/utils.ts";
import { ProjectStatus } from "../../types/project.ts";

export default function ProjectEdit() {
  const { id } = useParams();
  let projectNumericId: number = parseIdToNumber(id);
  const [userId, setUserId] = useState(0);
  const [roleType, setRoleType] = useState<number>(0);
  const [userNotInProjectList, setUserNotInProjectList] = useState<UserInfo[]>([]);
  const [projectInfo, setProjectInfo] = useState<ProjectInfo>({
    git_address: "",
    config: {
      join_by_self: false,
      must_check_by_other: false,
      secret: false
    }, description: "", id: 0, name: "", owner_id: 0, owner_name: "", status: ProjectStatus.INACTIVE
  });

  const navigate = useNavigate();
  let request = getRequestAndSetNavigateLocation(navigate, useLocation());

  const UserInProjectMethodsRef = useRef<UserInProjectMethods>(null);
  const messageContext = useContext(MessageContext);
  const { middleMessageApi } = messageContext as MessageContextValue;


  const handleAddUser = (newUser: UserInfo) => {
    UserInProjectMethodsRef.current?.addUser(newUser);
  };

  function addRole() {
    //userNotInProjectList 里 id是 userId的
    if (userId == 0 || roleType == 0) {
      return;
    }
    request.post("project/create_role", {
      project_id: projectNumericId,
      user_id: userId,
      role_type: roleType
    }).then((res) => {
      if (res.data.code === 0) {
        const user = userNotInProjectList.find(user => user.id === userId) !;
        user.role_type = roleType;
        handleAddUser(user);
        setUserNotInProjectList(prevList =>
          prevList.filter(user => user.id !== userId)
        );
        setUserId(0);
        setRoleType(0);
        middleMessageApi.success({
          content: "Role added successfully"
        }).then();
      } else {
        middleMessageApi.error({
          content: res.data.msg
        }).then();
      }
    }).catch((error) => {
      console.error("Error adding role:", error);
    });

  }

  function handleSave() {
      request.post("project/update", {
        config: projectInfo.config,
        name: projectInfo.name,
        id: projectNumericId,
        git_address: projectInfo.git_address,
        description: projectInfo.description,
      }).then((res) => {
        if (res.data.code === 0) {
          middleMessageApi.success({
            content: res.data.message
          }).then();
        } else {
          middleMessageApi.error({
            content: res.data.message
          }).then();
        }
      })
  }

  useEffect(() => {
    const fetchAllUsers = async () => {
      try {
        const allResponse = await request.get("user/list");
        if (allResponse.data.code === 0) {
          const allUsers = allResponse.data.data.list as UserInfo[];

          const inProjectResponse = await request.get(`project/user_list?id=${projectNumericId}`);
          if (inProjectResponse.data.code === 0) {
            const inProjectUsers = inProjectResponse.data.data.list as UserInfo[];

            // 计算 all - inProject
            const notInProjectUsers = allUsers.filter(user =>
              !inProjectUsers.some(inProjectUser => inProjectUser.id === user.id)
            );

            setUserNotInProjectList(notInProjectUsers);
          }
        }
      } catch (error) {
        console.error("Error fetching user list:", error);
      }
    };

    fetchAllUsers().then();
  }, []);

  useEffect(() => {
    request.get<BaseResponse<ProjectInfo>>(`project/info`, {
      params: {
        id: projectNumericId
      }
    }).then((res) => {
      if (res.data.code === 0) {
        setProjectInfo(res.data.data);
      } else {
        middleMessageApi.error({
          content: res.data.message
        }).then();
      }
    });
  }, []);



  return (
    <>
      <UserInProjectSelect id={projectNumericId} ref={UserInProjectMethodsRef} />

      <div
        style={{
          width: "80%", // 增大整个容器宽度
          margin: "0 auto", // 使容器居中
          backgroundColor: "#f0f0f0",
          padding: "1rem"
        }}
      >
        <Select
          placeholder="Select User"
          onChange={setUserId}
          onClear={() => setUserId(0)}
          style={{ width: 200 }}
          value={userId}
        >
          {userNotInProjectList.map((user) => (
            <Select.Option
              key={user.id}
              value={user.id}
              // selected={userId === user.id}
            >
              {user.user_name}
            </Select.Option>
          ))}
        </Select>

        <RoleTypeForAddSelect roleType={roleType} onChange={setRoleType} />
        <Button onClick={addRole}>Add Role</Button>
      </div>
      <div style={{
        textAlign: "center",
        marginTop: 100,
        display: "flex",
        flexDirection: "column",
        alignItems: "center"
      }}>
        <div className="divStyle">
          <label className="labelStyle">Name: </label>
          <Input className="inputStyle"
            onChange={(e: { target: { value: string; }; }) => setProjectInfo(prev => {
              let newProjectInfo = Object.assign({}, prev);
              newProjectInfo.name = e.target.value;
              return newProjectInfo;
            })} value={projectInfo?.name} />
        </div>
        <div className="divStyle">
          <label className="labelStyle">Description: </label>
          <Input.TextArea className="inputStyle"
            onChange={(e: { target: { value: string; }; }) => setProjectInfo(prev => {
              let newProjectInfo = Object.assign({}, prev);
              newProjectInfo.description = e.target.value;
              return newProjectInfo;
            })} value={projectInfo?.description} />
        </div>
        <div className="divStyle">
          <label className="labelStyle"> Git Local Path: </label>
          <Input className="inputStyle"
            onChange={(e: { target: { value: string; }; }) => setProjectInfo(prev => {
              let newProjectInfo = Object.assign({}, prev);
              newProjectInfo.git_address = e.target.value;
              return newProjectInfo;
            })} value={projectInfo?.git_address} />
        </div>
        <div className="divStyle">
          <Checkbox checked={projectInfo.config.must_check_by_other}
                   onChange={(e: { target: { checked: boolean; }; }) => setProjectInfo((prev) => {
                      let newProjectInfo = Object.assign({}, prev);
                      console.log(e.target.checked)
                      newProjectInfo.config.must_check_by_other = e.target.checked;
                      return newProjectInfo;
                   })}
                    value={projectInfo.config.must_check_by_other} > must_check_by_other </Checkbox>
          <Checkbox checked={projectInfo.config.join_by_self}
                    onChange={(e: { target: { checked: boolean; }; }) => setProjectInfo((prev) => {
                      let newProjectInfo = Object.assign({}, prev);
                      console.log(e.target.checked)
                      newProjectInfo.config.join_by_self = e.target.checked;
                      return newProjectInfo;
                    })}
                    value={projectInfo.config.join_by_self} > join_by_self </Checkbox>

        </div>
        <div className="divStyle">
          <Button type="primary" onClick={() => handleSave()}>保存</Button>
        </div>

      </div>
    </>
  );
}
