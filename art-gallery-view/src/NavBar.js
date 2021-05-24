import React from "react";
function NavBar() {
    return (
        <div className="flex flex-wrap items-center">
            <div className="flex flex-shrink md:w-1/3 justify-center md:justify-start text-black">Museum</div>

                <ul className="flex flex-row md:justify-end">
                    <li>Home</li>
                    <li>About</li>
                    <li>Contacts</li>
                </ul>

            </div>
        
    );
}
export default NavBar;
