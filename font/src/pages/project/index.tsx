import { useEffect, useState } from "react";
import { Table, Pagination, Space } from "antd";
import { useLocation, useNavigate } from "react-router-dom";
import getRequestAndSetNavigateLocaton from "../../services/axios.ts";
import ProjectStatusSelect from "../../components/projectStatus.tsx";
import { projectStatusMap, roleTypesMap } from "../../types/project.ts";
import { Project } from "../../types/response.ts";
import RoleTypeSelect from "../../components/roleType.tsx";
import { ownerOrAdmin } from "../../services/utils.ts";
import { useUserContext } from "../../context/userInfo.tsx";

const ProjectList = () => {


  const navigate = useNavigate();
  let request = getRequestAndSetNavigateLocaton(navigate, useLocation());

  const [currentPage, setCurrentPage] = useState(1);
  const [pageSize, setPageSize] = useState(5);

  const [roleType, setRoleType] = useState(0);
  const [selectedStatus, setSelectedStatus] = useState("");
  const [total, setTotal] = useState(0);
  const [ProjectList, setProjectList] = useState<Project[]>([]);

  const {user} = useUserContext()
  if (!user || Object.keys(user).length === 0) {
    return navigate("/");
  }

  const columns = [
    {
      title: "Name",
      dataIndex: "name",
      key: "name",
    },
    {
      title: "Status",
      dataIndex: "status",
      key: "status",
      render: (_text: any, record: Project) => {
        return projectStatusMap[record.status];
      },
    },
    {
      title: "Role Type",
      dataIndex: "role_type",
      key: "role_type",
      render: (_text: any, record: Project) => {
        return roleTypesMap[record.role_type];
      },
    },
    {
      title: "Owner Name",
      dataIndex: "owner_name",
      key: "owner_name",
    },
    {
      title: "Action",
      key: "action",
      render: (_text: any, record: Project) => (
        <Space size="middle">
          {ownerOrAdmin(user.id, record.owner_id) && (
            <a
              onClick={() => {
                navigate(`/project/edit/${record.id}`);
              }}
            >
              Edit
            </a>
          )}
          <a
            onClick={() => {
              navigate(`/project/tasks/${record.id}`);
            }}
          >
            Tasks
          </a>
          <a
            onClick={() => {
              navigate(`/project/${record.id}`);
            }}
          >
            View
          </a>
        </Space>
      ),
    },
  ];

  const fetchProjectList = () => {
    request
      .get("project/list", {
        params: {
          page: currentPage,
          page_size: pageSize,
          role_type: roleType,
          status: selectedStatus,
        },
      })
      .then((res) => {
        if (res.data.code == 0) {
          setTotal(res.data.data.total);
          setProjectList(res.data.data.list);
        }
      });
  };
  useEffect(fetchProjectList, [
    currentPage,
    pageSize,
    roleType,
    selectedStatus,
  ]);

  const handlePageChange = (page: number, size: number) => {
    setCurrentPage(page);
    setPageSize(size);
  };

  return (
    <div>
      <Space direction="horizontal" style={{ marginBottom: 16 }}>
        {user.id != 1 && (
          <RoleTypeSelect roleType={roleType} onChange={setRoleType} />
        )}
        <ProjectStatusSelect
          status={selectedStatus}
          onChange={setSelectedStatus}
        />
      </Space>
      <Pagination
        current={currentPage}
        pageSize={pageSize}
        total={total}
        onChange={handlePageChange}
        style={{ marginBottom: 16 }}
      />
      <Table
        dataSource={ProjectList}
        rowKey={(item) => item.id}
        columns={columns}
      />
      <Pagination
        current={currentPage}
        pageSize={pageSize}
        total={total}
        onChange={handlePageChange}
        showSizeChanger={true}
        style={{ marginTop: 16 }}
      />
    </div>
  );
};

export default ProjectList;
