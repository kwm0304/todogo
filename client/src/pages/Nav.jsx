import { Link } from 'react-router-dom';

const Nav = ({ isLoggedIn, username }) => {
  return (
    <nav className="bg-gradient-to-r from-teal-300 to-cyan-400 text-white py-4 px-8 flex items-center justify-between sticky top-0 z-10">
      <div className="text-xl font-bold">TODO</div>
      <div className="flex items-center">
        {isLoggedIn ? (
          <>
            <div className="mr-4 font-semibold">{username}</div>
            <Link to="/logout" className="mr-4 font-semibold">Logout</Link>
          </>
        ) : (
          <>
            <Link to="/signup" className="mr-4 font-semibold">Signup</Link>
            <Link to="/login" className='font-semibold'>Login</Link>
          </>
        )}
      </div>
    </nav>
  );
};

export default Nav;
