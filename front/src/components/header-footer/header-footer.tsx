import { Link, Outlet } from 'react-router-dom';
import { useState } from 'react';
import './header-footer.css';

export default function DeathNoteLayout() {
  const [menuOpen, setMenuOpen] = useState(false);

  return (
    <div className="deathnote-layout">
      <div className='dark-overlay'></div>
      <header className='dn-header'>
        <nav className='dn-navbar'>
          <Link to="/List" className="dn-logo">
            <h1 className='dn-title'>DeathNote</h1>
          </Link>
          <button className='dn-menu-toggle' onClick={() => setMenuOpen(!menuOpen)}>
            ☰
          </button>
          <div className={`dn-menu ${menuOpen ? 'open' : ''}`}>
            <Link to="/List" className="dn-link" onClick={() => setMenuOpen(false)}>Víctimas Registradas</Link>
            <Link to="/KillRegis" className="dn-link" onClick={() => setMenuOpen(false)}>Escribir un Nombre</Link>
          </div>
        </nav>
      </header>
      <main className='dn-main'>
        <Outlet />
      </main>
      <footer className='dn-footer'>
        <p>&copy; 2025 Programación Avanzada tercer corte</p>
      </footer>
    </div>
  );
}
