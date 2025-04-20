import { useNavigate } from 'react-router-dom';
import { Button } from "antd";

export  function NotFound  (){
  const navigate = useNavigate();
  const goHome = () => {
    navigate('/');
  };

  const goProject = () => {
    navigate('/project');
  };

  return (
    <div>
      <h1>404 - Page Not Found</h1>
      <Button onClick={goHome}>Go to Home</Button>
      <Button onClick={goProject}>Go to Project</Button>
    </div>
  );
}
