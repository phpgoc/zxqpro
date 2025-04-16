import { Pagination, Select, Space, Table } from "antd";
import {type Message} from "../../types/message.ts";
import { useContext, useEffect, useState } from "react";
import getRequestAndSetNavigateLocation from "../../services/axios.ts";
import { useLocation, useNavigate } from "react-router-dom";
import MessageContext, { type MessageContextValue, SetMessageNumberContext } from "../../context/message.tsx";
import CustomLink  from "../../components/customLink.tsx";

export default function MessageIndex() {
  const [total, setTotal] = useState(0);
  const [read, setRead] = useState(0);
  const [list, setList] = useState<Message[]>([]);
  const [page, setPage] = useState(1);
  const [pageSize, setPageSize] = useState(10);


  const navigate = useNavigate();
  let request = getRequestAndSetNavigateLocation(navigate, useLocation());

  const messageContext = useContext(MessageContext);
  const { middleMessageApi } = messageContext as MessageContextValue;
  const setMessageNumber = useContext(SetMessageNumberContext);

  function handleRead(id: number) {
    request.post("message/read", { id: id }).then((response) => {
      if (response.data.code == 0) {
        middleMessageApi.success(response.data.message).then();
        setMessageNumber((prev) => {
         return  prev - 1
        });
        setList((prev) => {
          return prev.map((item) => {
            if (item.id === id) {
              return { ...item, read: true };
            }
            return item;
          });
        })
      } else {
        middleMessageApi.warning(response.data.message).then();
      }
    }).then()
  }

  const columns = [
    {
      title: "From UserName",
      dataIndex: "user_name",
      key: "user_name"
    },
    {
      title: "Message",
      dataIndex: "message",
      key: "message"
    },
    {
      title: "Link",
      dataIndex: "link",
      key: "link",
      render: (_text: any, record: Message) => {
        if (record.link) {
          return (
            <CustomLink to={record.link} >
              {record.link}
            </CustomLink>
          );
        }
        return null;
      }
    },
    {
      title: "Time",
      dataIndex: "time",
      key: "time"
    },
    {
      title: "Action",
      key: "action",
      render: (_text: any, record: Message) => (
        <Space size="middle">
          {!record.read &&
            <a
              onClick={() => {
                handleRead(record.id);
              }}
            >
              Read
            </a>
          }
        </Space>
      )
    }
    ]


  function handlePageChange(page: number, pageSize: number) {
    setPage(page)
    setPageSize(pageSize)
  }

  useEffect(() => {
    request.get("message/receive_list",{
      params: {
        page: page,
        page_size: pageSize,
        read: read
      }
    }).then(
      (response) => {
        if (response.data.code == 0) {
          setList(response.data.data.list);
          setTotal(response.data.data.total);
        } else {
          middleMessageApi.error(response.data.message).then();
        }
      }
    )

  }, [read, page, pageSize]);

  return (
    <>
    <Space direction="horizontal" style={{ marginBottom: 16 }}>
      <Select
        placeholder="Read"
        value={read?"已读":"未读"}
        onChange={(value) => {
          setRead(value=="已读"?1:0);
        }}
        style={{ width: 200 }}
      >
        <Select.Option value={"已读"}>
            已读
        </Select.Option>
        <Select.Option value={"未读"}>
            未读
        </Select.Option>

      </Select>
    </Space>

      <div>
      <Pagination
        current={page}
        pageSize={pageSize}
        total={total}
        onChange={handlePageChange}
        style={{ marginBottom: 16 }}
      />
      <Table
        dataSource={list}
        rowKey={(item) => item.id}
        columns={columns}
      />
      <Pagination
        current={page}
        pageSize={pageSize}
        total={total}
        onChange={handlePageChange}
        showSizeChanger={true}
        style={{ marginTop: 16 }}
      />
      </div>
    </>
  )
}