import UserInProject from "../../components/userInProject.tsx";
import { useParams } from "react-router-dom";
import { parseIdToNumber } from "../../services/utils.ts";

export default function ProjectView() {
  const { id } = useParams();
  let numericId: number = parseIdToNumber(id);
  return (
    <>
      <UserInProject id={numericId} />
    </>
  )
}
