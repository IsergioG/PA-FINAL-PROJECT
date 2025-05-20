
import { Route, Routes } from 'react-router-dom'

import HeaderFooter from './components/header-footer/header-footer'
import KillRegister from './pages/kill-register/kill-register'
import DeathNoteNotFound from './pages/not-found/not-found'
import DeathNoteList from './pages/death-note-list/dn-list'


function App() {
  return (
    <Routes>  
      <Route path='/' element={<HeaderFooter/>}>
        <Route path='/List' element={<DeathNoteList/>}/>
        <Route path ='/KillRegis' element={<KillRegister/>}/>
        <Route path="*" element={<DeathNoteNotFound />} />
      </Route>
    </Routes>
  )
}

export default App
