import { BrowserRouter } from "react-router-dom";
import "./App.css";
import { Container } from "./styled";
import { RoutesManagement } from "./routes/Routes";

function App() {
  return (
    <Container>
      <BrowserRouter>
        <RoutesManagement />
      </BrowserRouter>
    </Container>
  );
}

export default App;
