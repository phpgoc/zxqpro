import React, { useState } from 'react';
import { Input, Select, Table, Pagination, Space } from 'antd';
import {useNavigate} from "react-router-dom";
import getRequestAndSetNavigate from "../../services/axios.ts";
const { Search } = Input;
const { Option } = Select;

const ProjectList = () => {
    const [currentPage, setCurrentPage] = useState(1);
    const [pageSize, setPageSize] = useState(10);
    const [searchText, setSearchText] = useState('');
    const [selectedStatus, setSelectedStatus] = useState('');
    const navigate = useNavigate();
    let request = getRequestAndSetNavigate(navigate);

    const dataSource = [
        {
            key: '1',
            name: 'Project 1',
            status: 'Active',
        },
        {
            key: '2',
            name: 'Project 2',
            status: 'Inactive',
        },
        // 可以添加更多项目数据
    ];

    const columns = [
        {
            title: 'Name',
            dataIndex: 'name',
            key: 'name',
        },
        {
            title: 'Status',
            dataIndex: 'status',
            key: 'status',
        },
    ];

    const handleSearch = (value) => {
        setSearchText(value);
    };

    const handleStatusChange = (value) => {
        setSelectedStatus(value);
    };

    const handlePageChange = (page, size) => {
        setCurrentPage(page);
        setPageSize(size);
    };

    // 过滤数据
    const filteredData = dataSource.filter((item) => {
        const matchSearch = item.name.includes(searchText);
        const matchStatus = selectedStatus === '' || item.status === selectedStatus;
        return matchSearch && matchStatus;
    });

    return (
        <div>
            <Space direction="horizontal" style={{ marginBottom: 16 }}>
                <Search
                    placeholder="Search by project name"
                    onSearch={handleSearch}
                    style={{ width: 200 }}
                />
                <Select
                    placeholder="Select status"
                    onChange={handleStatusChange}
                    style={{ width: 200 }}
                >
                    <Option value="Active">Active</Option>
                    <Option value="Inactive">Inactive</Option>
                </Select>
            </Space>
            <Pagination
                current={currentPage}
                pageSize={pageSize}
                total={filteredData.length}
                onChange={handlePageChange}
                style={{ marginBottom: 16 }}
            />
            <Table
                dataSource={filteredData.slice((currentPage - 1) * pageSize, currentPage * pageSize)}
                columns={columns}
            />
            <Pagination
                current={currentPage}
                pageSize={pageSize}
                total={filteredData.length}
                onChange={handlePageChange}
                style={{ marginTop: 16 }}
            />
        </div>
    );
};

export default ProjectList;
    