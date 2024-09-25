import React, {useState} from "react";
import {DingdingOutlined} from "@ant-design/icons";
import {Button} from "antd";
import {AuthService} from "@/services";

const DingTalk: React.FC<{ providers: string[] }> = (props) => {
    const [loading, setLoading] = useState(false)
    const {providers} = props;
    if (!providers.includes("dingtalk")) {
        return null;
    }
    return (
        <Button size={"large"} loading={loading} onClick={() => {
            setLoading(true);
            AuthService.getAuthorizeUrl("dingtalk").then(({authorizeUrl}) => {
                window.location.href = authorizeUrl
            }).catch(() => {
                setLoading(false)
            })
        }} icon={<DingdingOutlined/>}>
            DingTalk
        </Button>
    )
}

export default DingTalk;
