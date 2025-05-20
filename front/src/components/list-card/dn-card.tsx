import { useState } from "react";
import "./dn-card.css";
import type { Kill } from "../../types/Kill";

export default function VictimCard({ kill }: { kill: Kill }) {
  const [showPopup, setShowPopup] = useState(false);

  return (
    <>
      <div className="dn-card" onClick={() => setShowPopup(true)}>
        <h2 className="dn-card-name">{kill.fullName}</h2>
        <p className="dn-card-cause"><strong>Causa:</strong> {kill.causeOfDeath || "Ataque al corazón"}</p>
        <p className="dn-card-hint">Haz clic para ver los detalles de la muerte</p>
      </div>

      {showPopup && (
        <div className="dn-popup-overlay" onClick={() => setShowPopup(false)}>
          <div className="dn-popup-content" onClick={(e) => e.stopPropagation()}>
            <h3 className="dn-popup-title">Detalles de la Muerte</h3>
            <p className="dn-popup-text">{kill.details || "Sin detalles."}</p>
            <button className="dn-popup-close" onClick={() => setShowPopup(false)}>Cerrar Cuaderno</button>
          </div>
        </div>
      )}
    </>
  );
}
