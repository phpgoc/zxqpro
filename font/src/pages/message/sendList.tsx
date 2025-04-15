import { Pagination,  Space, Table } from "antd";
import {type Message} from "../../types/message.ts";
import { useContext, useEffect, useState } from "react";
import getRequestAndSetNavigateLocation from "../../services/axios.ts";
import { useLocation, useNavigate } from "react-router-dom";
import MessageContext, { type MessageContextValue } from "../../context/message.tsx";

export default function MessageSend() {
  const [total, setTotal] = useState(0);
  const [list, setList] = useState<Message[]>([]);
  const [page, setPage] = useState(1);
  const [pageSize, setPageSize] = useState(10);


  const navigate = useNavigate();
  let request = getRequestAndSetNavigateLocation(navigate, useLocation());

  const messageContext = useContext(MessageContext);
  const { middleMessageApi } = messageContext as MessageContextValue

  function handleLink(link: string, newTab: boolean = false) {
      if (newTab){
        window.open(link, "_blank");
      }else{
        window.location.href = link
      }
  }

  const columns = [
    {
      title: "To UserName",
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
      key: "link"
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
          {record.link && (
            <a
              onClick={() => {
                handleLink(record.link!);
              }}
            >
              Go to Link
            </a>
          )}
          {record.link && (
            <a
              onClick={() => {
                handleLink(record.link!, true);
              }}
            >
              Open Link with New Tab
            </a>
          )}
        </Space>
      )
    }
  ]


  function handlePageChange(page: number, pageSize: number) {
    setPage(page)
    setPageSize(pageSize)
  }

  useEffect(() => {
    request.get("message/send_list",{
      params: {
        page: page,
        page_size: pageSize,
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

  }, [ page, pageSize]);

  return (
    <>

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