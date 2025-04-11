import {useContext, useState} from "react";
import getRequestAndSetNavigate from "../../services/axios.ts";
import {useNavigate} from "react-router-dom";
import {Input, Button} from "antd";
import MessageContext,{type MessageContextValue} from "../../context/message.tsx";
import {BaseResponse} from "../../types/response.ts";


export default function Register() {
    const [name, setName] = useState("");
    const [password, setPassword] = useState("");
    const navigate = useNavigate();
    let request = getRequestAndSetNavigate(navigate);
    const messageContext = useContext(MessageContext)
    const {middleApi} = messageContext as MessageContextValue;

    const handleSubmit = (e: React.FormEvent<HTMLFormElement>) => {
        e.preventDefault();
        request.post<BaseResponse>("/admin/register", {name, password})
            .then((response) => {
                if (response.data.code== 0) {
                    middleApi.success(response.data.message).then((_: any) => {
                        setName("")
                        setPassword("")
                    });
                }else {
                     middleApi.warning(response.data.message).then()
                }
            })
            .catch((_) => {
                middleApi.error("Registration failed. Please try again.").then()
            });
    };

    return (
        <form onSubmit={handleSubmit}>

            <div style={{ textAlign: 'center', marginBottom: 32 }}>
            <Input
                value={name}
                onChange={(e) => setName(e.target.value)}
                placeholder="Name"
                style={{ width: '50%', margin: '20px 25%' }}
            />
            <br />
            <Input.Password
                value={password}
                onChange={(e) => setPassword(e.target.value)}
                placeholder="Password"
                style={{ width: '50%', margin: '20px 25%' }}
            />
            <br />

            <Button type="primary" htmlType="submit" style={{ width: '15%',margin: '20px 5%' }}>
                Register
            </Button>
            <Button  onClick={() => {setName(""); setPassword("");}} style={{ width: '15%',margin: '20px 5%' }}>
                Reset
            </Button>
            </div>

        </form>
    );
}