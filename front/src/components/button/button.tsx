import React from 'react';
import './button.css';

type ButtonSize = 'small' | 'medium' | 'large';
type ButtonVariant = 'primary' | 'secondary';
type ButtonType = 'button' | 'submit' | 'reset';

interface ButtonProps {
  label: string;
  onClick?: () => void;
  size?: ButtonSize;
  variant?: ButtonVariant;
  className?: string;
  disabled?: boolean;
  type?: ButtonType;
}

const DeathNoteButton: React.FC<ButtonProps> = ({
  label,
  onClick,
  size = 'medium',
  variant = 'primary',
  className = '',
  disabled = false,
  type = 'button'
}) => {
  return (
    <button
      type={type}
      disabled={disabled}
      onClick={onClick}
      className={`dn-btn dn-${size} dn-${variant} ${className}`}
    >
      {label}
    </button>
  );
};

export default DeathNoteButton;
