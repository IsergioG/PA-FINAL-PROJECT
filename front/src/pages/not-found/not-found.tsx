import { Link } from "react-router-dom";
import Button from "../../components/button/button";
import "./not-found.css";

export default function DeathNoteNotFound() {
  return (
    <div className="dn-notfound-container">
      <h1 className="dn-notfound-title">404 - Página del Shinigami no encontrada</h1>
      <p className="dn-notfound-text">
        Este nombre no está registrado... o la página no existe.
      </p>
      <Link to="/List">
        <Button
          label="Volver al Cuaderno"
          size="small"
          variant="primary"
        />
      </Link>
    </div>
  );
}
