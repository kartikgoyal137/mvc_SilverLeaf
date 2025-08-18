import { Navigate } from 'react-router-dom';
import { getUser } from '../utils/Auth';

export default function ProtectedRoute({ children, allowedRoles }) {
  const user = getUser();

  if (!user) {
    return <Navigate to="/login" replace />;
  }

  if (allowedRoles && !allowedRoles.includes(user.role)) {
    return <h1>Access Denied</h1>;
  }

  return children;
}
