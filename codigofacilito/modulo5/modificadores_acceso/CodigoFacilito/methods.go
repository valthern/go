package modificadores

// Método Privado, debido a que empieza su nombre con minúscula
func (c *Curso) getTitulo() string {
	return c.Titulo
}

// Método Público, debido a que su nombre empieza con mayúscula
func (self *Curso) ObtenerTitulo() string {
	return self.getTitulo()
}
