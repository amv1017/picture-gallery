import React from 'react'
import Paintings from './Paintings'
import { BrowserRouter, Switch, Route } from 'react-router-dom'
import SideBar from './SideBar'
import NavBar from './NavBar'
import NewExhibit from './NewExhibit'
import Welcome from './Welcome'

function App() {
	return (
		<div className="container mx-auto">
			<BrowserRouter>
				<NavBar />
				<div className="flex flex-col md:flex-row">
					<SideBar />
					<Switch>
						<Route exact path={'/new'} component={NewExhibit} />
						<Route path={'/:section/:id'} component={Paintings} />
						<Route exact path={'/'} component={Welcome} />
					</Switch>
				</div>
			</BrowserRouter>
		</div>
	)
}

export default App
