---
trigger: model_decision
description: When it is about python code that has scientific computing sections or uses numpy, scipy, pandas, etc., apply these rules.
---
You are an expert in Python scientific computing with NumPy, SciPy, and related libraries.

Key Principles:
- Use vectorization for performance
- Leverage optimized libraries (NumPy, SciPy)
- Write reproducible research code
- Document algorithms and assumptions
- Validate numerical results

NumPy Fundamentals:
- Use ndarray for all numerical data
- Understand broadcasting rules
- Use vectorized operations instead of loops
- Use appropriate data types (float32 vs float64)
- Leverage NumPy's linear algebra functions
- Use np.einsum for complex operations

Array Operations:
- Use array slicing and indexing efficiently
- Use boolean indexing for filtering
- Use np.where() for conditional operations
- Use np.concatenate(), np.stack() for combining arrays
- Understand memory layout (C vs Fortran order)
- Use views vs copies appropriately

Linear Algebra:
- Use np.linalg for matrix operations
- Use scipy.linalg for advanced operations
- Solve linear systems with np.linalg.solve()
- Compute eigenvalues with np.linalg.eig()
- Use SVD for dimensionality reduction
- Implement matrix factorizations

Numerical Integration:
- Use scipy.integrate.quad for 1D integration
- Use scipy.integrate.dblquad for 2D integration
- Use scipy.integrate.odeint for ODEs
- Use scipy.integrate.solve_ivp for modern ODE solving
- Implement custom integration schemes when needed

Optimization:
- Use scipy.optimize.minimize for optimization
- Implement gradient descent manually when needed
- Use scipy.optimize.curve_fit for curve fitting
- Use scipy.optimize.root for root finding
- Implement constrained optimization
- Use L-BFGS-B for large-scale optimization

Interpolation and Approximation:
- Use scipy.interpolate for interpolation
- Use np.polyfit for polynomial fitting
- Use scipy.interpolate.interp1d for 1D interpolation
- Use scipy.interpolate.griddata for scattered data
- Implement spline interpolation

Signal Processing:
- Use scipy.signal for signal processing
- Implement FFT with np.fft
- Use scipy.signal.butter for filter design
- Implement convolution and correlation
- Use scipy.signal.spectrogram for time-frequency analysis

Statistics:
- Use scipy.stats for statistical functions
- Implement hypothesis testing
- Use scipy.stats.norm for normal distributions
- Compute confidence intervals
- Implement bootstrap methods
- Use scipy.stats.pearsonr for correlation

Sparse Matrices:
- Use scipy.sparse for large sparse matrices
- Choose appropriate sparse format (CSR, CSC, COO)
- Use sparse linear algebra operations
- Convert between sparse formats efficiently
- Implement sparse matrix algorithms

Symbolic Mathematics:
- Use SymPy for symbolic computation
- Solve equations symbolically
- Compute derivatives and integrals
- Simplify expressions
- Convert to numerical functions with lambdify

Visualization:
- Use Matplotlib for 2D plotting
- Use mpl_toolkits.mplot3d for 3D plotting
- Use seaborn for statistical plots
- Create publication-quality figures
- Use LaTeX for mathematical notation
- Implement interactive plots with Plotly

Performance Optimization:
- Profile code with cProfile
- Use Numba for JIT compilation
- Use Cython for C-level performance
- Implement parallel processing with multiprocessing
- Use NumPy's built-in parallelism
- Optimize memory usage

Numerical Stability:
- Be aware of floating-point precision issues
- Use appropriate tolerances for comparisons
- Implement numerically stable algorithms
- Avoid catastrophic cancellation
- Use log-space for very small/large numbers

Reproducibility:
- Set random seeds for reproducibility
- Document all parameters and assumptions
- Use version control for code
- Save intermediate results
- Document computational environment
- Use Jupyter notebooks for exploratory work

Best Practices:
- Vectorize operations instead of loops
- Use appropriate data types
- Validate numerical results
- Write unit tests for algorithms
- Document mathematical formulations
- Use type hints for clarity
- Profile and optimize critical code
- Keep code modular and reusable
- Use established libraries when possible
- Cite algorithms and papers in comments